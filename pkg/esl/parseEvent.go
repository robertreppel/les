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
			var existingItemIndex = foundPreviously(lineItems, eventName)
			if len(propertiesList) > 0 {
				inputProperties := strings.Split(propertiesList, ",")
				for _, inputProperty := range inputProperties {
					var propertyName =strings.Trim(inputProperty, " ")
					propertyName = strings.Split(propertyName, "=")[0] // Ignore example values, e.g. email=asdf@gmail.com
					var parsedProperty = Property{Name: propertyName}
					properties = append(properties, parsedProperty)
				}
				if existingItemIndex != -1 {
					lineItems[existingItemIndex] = Event{Name: eventName, Properties: properties}
				} else {
					lineItems = append(lineItems, Event{Name: eventName, Properties: properties})
				}

			} else {
				if existingItemIndex == -1 {
					lineItems = append(lineItems, Event{Name: eventName})
				}
			}

		}
	} else {
		var eventName = strings.Trim(eslInput, " ")
		var itemIndex = foundPreviously(lineItems, eventName)
		if itemIndex != -1 {
			return lineItems
		}
		lineItems = append(lineItems, Event{Name: strings.Trim(eslInput, " ")})
	}
	return lineItems
}

func foundPreviously(lineItems []Item, itemName string) int {
	for index, item := range lineItems {
		switch item.(type) {
		case Event:
			if item.(Event).Name == itemName {
				return index
			}
		case Command:
			if item.(Command).Name == itemName {
				return index
			}
		case Document:
			if item.(Document).Name == itemName {
				return index
			}
		}
	}
	return -1
}
