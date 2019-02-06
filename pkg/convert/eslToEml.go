package convert

import (
	"fmt"
	"strings"
	"unicode"

	"github.com/robertreppel/les/pkg/eml"
	"github.com/robertreppel/les/pkg/esl"
)

// EslToEml converts an event markdown (esl) spec to event markup language (eml)
// - Event Storming Language is the format used on (e.g.) webeventstorming.com
// - Event Modeling Language is used to specify event sourced systems for the Adaptech Platform.
func EslToEml(markdown esl.Esl) (EslToEmlConversion, error) {
	result := EslToEmlConversion{
		Esl: eml.Solution{},
		MarkdownValidationErrors: []esl.EslValidationError{},
	}

	boundedContext := getBoundedContext(markdown)
	result.Esl.EmlVersion = "0.1-alpha"
	result.Esl.Name = boundedContext.Name // ESL can only have one bounded context right now. Use it's name for the solution name.
	contexts := []eml.BoundedContext{}

	contexts = append(contexts, boundedContext)
	result.Esl.Contexts = contexts

	result = getStreams(markdown, result)

	// Build an index which gets a list of all events which have a given property, across all bounded contexts.
	// This is so that SubscribesTo events for populating read models can be determined based on property naming conventions:
	// All events which have a property needed in the read model are going to be subscribed to.
	streamIDs := make(map[string]bool)
	eventsByPropertyLookup := make(map[string][]string)
	for _, context := range result.Esl.Contexts {
		for _, stream := range context.Streams {
			streamIDs[strings.ToLower(stream.Name+"id")] = true
			for _, event := range stream.Events {
				for _, property := range event.Event.Properties {
					if contains(event.Event.Name, eventsByPropertyLookup[property.Name]) == false {
						eventsByPropertyLookup[property.Name] = append(eventsByPropertyLookup[property.Name], event.Event.Name)
					}
				}
			}
		}
	}

	readmodels, validationErrors := getReadmodels(markdown, eventsByPropertyLookup, streamIDs)
	for _, eachError := range validationErrors {
		result.MarkdownValidationErrors = append(result.MarkdownValidationErrors, eachError)
	}

	result.Esl.Contexts[0].Readmodels = readmodels

	result.Esl.Validate()
	return result, nil
}

func contains(eventID string, eventList []string) bool {
	hasValue := false
	for _, value := range eventList {
		if value == eventID {
			hasValue = true
		}
	}
	return hasValue
}

func getBoundedContext(markdown esl.Esl) eml.BoundedContext {
	for _, item := range markdown.Lines {
		switch item.(type) {
		case esl.Comment:
			boundedContext := item.(esl.Comment).Text
			return eml.BoundedContext{Name: boundedContext}
		}
	}
	return eml.BoundedContext{Name: ""}
}

func getReadmodels(markdown esl.Esl, eventsByPropertyLookup map[string][]string, streams map[string]bool) ([]eml.Readmodel, []esl.EslValidationError) {
	validationErrors := []esl.EslValidationError{}
	models := []eml.Readmodel{}
	for _, item := range markdown.Lines {
		switch item.(type) {
		case esl.Document:
			readmodelEsl := item.(esl.Document)
			readmodelEml := eml.Readmodel{}
			readmodelEml.Readmodel.Name = strings.Replace(readmodelEsl.Name, " ", "", -1)
			properties := []eml.Property{}
			for _, eslProperty := range readmodelEsl.Properties {
				emlProperty := eml.Property{Name: strings.Replace(eslProperty.Name, " ", "", -1)}
				properties = append(properties, emlProperty)
			}
			if len(properties) > 0 {
				// The first property in the ESL document is assumed to be the key.
				// e.g. 'User List* // userId, name, email'
				// results in readmodel key = userId
				//Ensures consistent casing for stream Ids: e.g. TimesheetHoursId is represented as "timesheethoursId".
				streamName := properties[0].Name[:len(properties[0].Name)-2]
				key := firstCharToLower(streamName) + "Id"
				readmodelEml.Readmodel.Key = key
			}

			// Determine the events the read model needs to subscribe to, based on which events contain the properties of the read model:
			subscribesTo := make(map[string]bool)
			for _, property := range properties {
				_, isStreamID := streams[strings.ToLower(property.Name)]
				if !isStreamID {
					if eventList, ok := eventsByPropertyLookup[property.Name]; ok {
						for _, event := range eventList {
							subscribesTo[event] = true
						}
					} else {
						validationError := esl.EslValidationError{
							ErrorID: "MissingSubscribesToEventForReadmodelProperty",
							Message: "Could not find an event to subscribe to for getting readmodel property '" + property.Name + "'",
						}
						validationErrors = append(validationErrors, validationError)
					}
				}
			}
			subscribesToEventList := []string{}
			for k := range subscribesTo {
				subscribesToEventList = append(subscribesToEventList, k)
			}
			readmodelEml.Readmodel.SubscribesTo = subscribesToEventList
			models = append(models, readmodelEml)
		}
	}

	return models, validationErrors
}

func getStreams(markdown esl.Esl, result EslToEmlConversion) EslToEmlConversion {
	validationErrors := []esl.EslValidationError{}
	foundEvents := make(map[string][]eml.Event)
	foundCommands := make(map[string][]eml.Command)
	var mostRecentCommand *eml.Command

	for _, item := range markdown.Lines {
		switch item.(type) {
		case esl.Command:
			eslCommand := item.(esl.Command)
			commandID := strings.Replace(eslCommand.Name, " ", "", -1)
			newEmlCommand := eml.Command{}
			newEmlCommand.Command.Name = commandID

			emlParameters := []eml.Parameter{}
			for _, eslParameter := range eslCommand.Parameters {
				parameterNameWithoutSpaces := strings.Replace(eslParameter.Name, " ", "", -1)
				emlParameter := eml.Parameter{Name: parameterNameWithoutSpaces, Type: "string"}
				emlParameters = append(emlParameters, emlParameter)
			}
			newEmlCommand.Command.Parameters = emlParameters

			mostRecentCommand = &newEmlCommand

		case esl.Event:
			eslEvent := item.(esl.Event)
			eventName := eslEvent.Name
			eventNameWords := strings.Split(eventName, " ")
			if firstWordIsStreamName(eventNameWords) {
				firstWordInEventName := eventNameWords[0]
				streamName := firstWordInEventName
				newEmlEvent := eml.Event{}
				newEmlEvent.Event.Name = strings.Replace(eventName, " ", "", -1)
				emlProperties := []eml.Property{}
				for _, eslProperty := range eslEvent.Properties {
					propertyNameWithoutSpaces := strings.Replace(eslProperty.Name, " ", "", -1)
					emlProperty := eml.Property{Name: propertyNameWithoutSpaces, Type: "string"}
					if eslProperty.Name == "password" {
						emlProperty.IsHashed = true
					}
					emlProperties = append(emlProperties, emlProperty)
				}

				newEmlEvent.Event.Properties = emlProperties

				newEmlEvent = ensureThatEventHasAggregateIDProperty(newEmlEvent, streamName)

				if events, ok := foundEvents[streamName]; ok {
					events := append(events, newEmlEvent)
					foundEvents[streamName] = events
				} else {
					events := []eml.Event{newEmlEvent}
					foundEvents[streamName] = events
				}

				if mostRecentCommand != nil {
					mostRecentCommand = ensureThatCommandHasAggregateIDParameter(mostRecentCommand, streamName)
					mostRecentCommand = ensureThatAggregateIDParameterIsRequiredField(mostRecentCommand, streamName)

					if commands, ok := foundCommands[streamName]; ok {
						commands := append(commands, *mostRecentCommand)
						foundCommands[streamName] = commands
					} else {
						commands := []eml.Command{*mostRecentCommand}
						foundCommands[streamName] = commands
					}
					mostRecentCommand = nil
				}

			} else {
				validationError := esl.EslValidationError{ErrorID: "NoStreamName", Message: "'" + eventName + "': Could not determine event stream name."}
				validationErrors = append(validationErrors, validationError)
			}
		}
	}

	commandPostconditions := getCommandPostconditions(markdown)
	streams := []eml.Stream{}
	for streamName := range foundEvents {
		commandsWithPostconditions := []eml.Command{}
		commands := foundCommands[streamName]
		for _, command := range commands {
			postConditionEventIDs := commandPostconditions[streamName+"."+command.Command.Name]
			command.Command.Postconditions = postConditionEventIDs
			commandsWithPostconditions = append(commandsWithPostconditions, command)
		}
		stream := eml.Stream{
			Name:     streamName,
			Events:   foundEvents[streamName],
			Commands: commandsWithPostconditions,
		}
		streams = append(streams, stream)
	}
	result.Esl.Contexts[0].Streams = streams
	result.MarkdownValidationErrors = validationErrors
	return result
}

func firstWordIsStreamName(eventNameWords []string) bool {
	return len(eventNameWords) > 1
}

func getCommandPostconditions(markdown esl.Esl) map[string][]string {
	postConditions := make(map[string][]string)
	mostRecentCommandID := ""
	for _, item := range markdown.Lines {
		switch item.(type) {
		case esl.Command:
			eslCommand := item.(esl.Command)
			mostRecentCommandID = strings.Replace(eslCommand.Name, " ", "", -1)

		case esl.Event:
			// If there are events to the left of the first command in the esl event storming, these
			// are assumed to be external events. They are being ignored, for now. Need conventions
			// for integration between bounded contexts later.
			if mostRecentCommandID != "" {

				eslEvent := item.(esl.Event)
				eventName := eslEvent.Name
				eventNameWords := strings.Split(eventName, " ")
				firstWordInEventName := eventNameWords[0]
				streamName := firstWordInEventName
				eventID := strings.Replace(eventName, " ", "", -1)
				postConditions[streamName+"."+mostRecentCommandID] = append(postConditions[streamName+"."+mostRecentCommandID], eventID)
			}
		}
	}
	return postConditions
}

func ruleExists(rule string, rules []string) bool {
	for _, existingRule := range rules {
		if existingRule == rule {
			return true
		}
	}
	return false
}

func ensureThatCommandHasAggregateIDParameter(command *eml.Command, streamName string) *eml.Command {
	// Ensure that the command has a mandatory aggregate id parameter:
	hasAggregateID := false
	for index, parameter := range command.Command.Parameters {
		if parameter.Name == firstCharToLower(streamName)+"Id" {
			//Ensure consistent casing for stream Ids: e.g. TimesheetHoursId is represented as "timesheetHoursId".
			command.Command.Parameters[index].Name = firstCharToLower(streamName) + "Id"
			if !ruleExists("IsRequired", command.Command.Parameters[index].Rules) {
				command.Command.Parameters[index].Rules = append(command.Command.Parameters[index].Rules, "IsRequired")
			}
			hasAggregateID = true
		}
	}
	if !hasAggregateID {
		parameterValidationRules := []string{"IsRequired"}
		aggregateIDParameter := eml.Parameter{Name: firstCharToLower(streamName) + "Id", Rules: parameterValidationRules, Type: "string"}
		command.Command.Parameters = append(command.Command.Parameters, aggregateIDParameter)
	}
	return command
}

func ensureThatEventHasAggregateIDProperty(event eml.Event, streamName string) eml.Event {
	hasAggregateID := false
	for _, parameter := range event.Event.Properties {
		if parameter.Name == firstCharToLower(streamName+"Id") {
			hasAggregateID = true
			break
		}
	}

	var aggregateIDParameter eml.Property
	if !hasAggregateID {
		aggregateIDParameter = eml.Property{Name: firstCharToLower(streamName) + "Id", Type: "string"}
		event.Event.Properties = append(event.Event.Properties, aggregateIDParameter)
	}
	return event
}

func firstCharToLower(s string) string {
	if len(s) == 0 {
		return s
	}
	a := []rune(s)
	a[0] = unicode.ToLower(a[0])
	return (string(a))
}

func ensureThatAggregateIDParameterIsRequiredField(command *eml.Command, streamName string) *eml.Command {
	parameters := []eml.Parameter{}
	for _, parameter := range command.Command.Parameters {
		if strings.HasSuffix(parameter.Name, "Id") {
			if parameter.Name == firstCharToLower(streamName)+"Id" {
				if !parameter.RuleExists("IsRequired") {
					parameter.Rules = append(parameter.Rules, "IsRequired")
				}
			} else {
				if !parameter.RuleExists("IsRequired") {
					referencedStreamIDPrefix := parameter.Name[firstUppercaseLetterIn(parameter.Name):][:len(parameter.Name[firstUppercaseLetterIn(parameter.Name):])-2]
					if referencedStreamIDPrefix == "" {
						referencedStreamIDPrefix = parameter.Name[:len(parameter.Name)-2]
					}
					readModel := strings.Title(referencedStreamIDPrefix) + "Lookup"
					mustExist := "MustExistIn " + readModel
					parameter.Rules = append(parameter.Rules, mustExist)
				}
			}
		}
		parameters = append(parameters, parameter)
	}
	command.Command.Parameters = parameters
	return command
}

// EslToEmlConversion result
type EslToEmlConversion struct {
	Esl                      eml.Solution
	MarkdownValidationErrors []esl.EslValidationError
}

func firstUppercaseLetterIn(s string) int {
	for i, c := range s {
		char := fmt.Sprintf("%c", c)
		if char == strings.ToUpper(char) {
			return i
		}
	}
	return -1
}
