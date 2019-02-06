package esl

import (
	"regexp"
	"strings"
)

func parseDocument(eslInput string, lineItems []Item) []Item {
	if strings.Contains(eslInput, ":") {
		re, err := regexp.Compile("^(.*) *\\* *\\: *(.*)")
		if err != nil {
			panic(err)
		}
		document := re.FindAllStringSubmatch(eslInput, -1)
		if len(document) > 0 {
			var properties []Property
			first := document[0]
			propertiesList := first[2]
			propertiesList = strings.Trim(propertiesList, ", ")
			inputProperties := strings.Split(propertiesList, ",")
			for _, inputParameter := range inputProperties {
				if inputParameter != "" {
					var parsedParameter = Property{Name: strings.Trim(inputParameter, " ")}
					properties = append(properties, parsedParameter)
				}
			}
			lineItems = append(lineItems, Document{Name: strings.Trim(document[0][1], " "), Properties: properties})
			return lineItems
		}
	} else {
		spacesRemoved := strings.Trim(eslInput, " ")
		documentName := strings.Replace(spacesRemoved, "*", "", -1)
		lineItems = append(lineItems, Document{Name: documentName})
	}
	return lineItems
}