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
			var eventName = strings.Trim(event[0][1], " ")
			if len(propertiesList) > 0 {
				inputProperties := strings.Split(propertiesList, ",")
				for _, inputProperty := range inputProperties {
					var propertyName = strings.Trim(inputProperty, " ")
					propertyName = strings.Split(propertyName, "=")[0] // Ignore example values, e.g. email=asdf@gmail.com
					var parsedProperty = Property{Name: propertyName}
					properties = append(properties, parsedProperty)
				}
				lineItems = append(lineItems, Event{Name: eventName, Type: "Event", Properties: properties})

			} else {
				lineItems = append(lineItems, Event{Name: eventName, Type: "Event"})
			}
		}
	} else {
		lineItems = append(lineItems, Event{Name: strings.Trim(eslInput, " "), Type: "Event"})
	}
	return lineItems
}
