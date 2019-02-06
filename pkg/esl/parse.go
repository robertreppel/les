package esl

import (
	"strings"
)

// Parse Event Storming Language to Event Modeling Language Language
func Parse(eslInput []string) (Esl, error) {
	eventMarkdown := Esl{}
	var lineItems = []Item{}
	for _, line := range eslInput {
		if len(line) <= 2 || line[:2] == "//" {
			continue
		}
		if strings.Contains(line, "#") {
			lineItems = parseComment(line, lineItems)
		} else if !strings.Contains(line, "->") && !strings.Contains(line, "*") {
			lineItems = parseEvent(line, lineItems)
		} else if strings.Contains(line, "->") {
			lineItems = parseCommand(line, lineItems)
		} else if strings.Contains(line, "*") {
			lineItems = parseDocument(line, lineItems)
		}
	}
	eventMarkdown.Lines = lineItems
	return eventMarkdown, nil
}
