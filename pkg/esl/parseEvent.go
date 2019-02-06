package esl

import (
	"regexp"
	"strings"
)

func parseEvent(eslInput string, lineItems []Item) []Item {
	if strings.Contains(eslInput, ":") {
		re, err := regexp.Compile("^(.*) *\\: *(.*)")
		if err != nil {
			panic(err)
		}
		event := re.FindAllStringSubmatch(eslInput, -1)
		if len(event) > 0 {
			var properties []Property
			first := event[0]
			propertiesList := first[2]
			propertiesList = strings.Trim(propertiesList, ", ")
			if len(propertiesList) > 0 {
				inputProperties := strings.Split(propertiesList, ",")
				for _, inputProperty := range inputProperties {
					var parsedProperty = Property{Name: strings.Trim(inputProperty, " ")}
					properties = append(properties, parsedProperty)
				}
				lineItems = append(lineItems, Event{Name: strings.Trim(event[0][1], " "), Properties: properties})
			} else {
				lineItems = append(lineItems, Event{Name: strings.Trim(event[0][1], " ")})
			}

		}
	} else {
		lineItems = append(lineItems, Event{Name: strings.Trim(eslInput, " ")})
	}
	return lineItems
}