package esl

import (
	"strings"
)

func parseCommand(eslInput string, lineItems []Item) []Item {
	eslInput = strings.Trim(eslInput, " ")
	parts := strings.Split(eslInput, "->")
	commandName := strings.Trim(parts[0], " ")
	propertiesString := strings.Replace(parts[1], "/", "", -1)
	propertiesString = strings.Replace(propertiesString, " ", "", -1)
	inputParameters := strings.Split(propertiesString, ",")
	var parameters []Parameter
	for _, inputParameter := range inputParameters {
		if len(inputParameter) > 0 {
			var parsedParameter = Parameter{Name: inputParameter}
			parameters = append(parameters, parsedParameter)
		}
	}
	lineItems = append(lineItems, Command{Name: commandName, Parameters: parameters})
	return lineItems
}
