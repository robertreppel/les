package esl

import (
	"regexp"
	"strings"
)

func parseComment(eslInput string, lineItems []Item) []Item {
	re, _ := regexp.Compile("^ *# *(.*)")
	comment := re.FindAllStringSubmatch(eslInput, -1)
	if len(comment) > 0 {
		lineItems = append(lineItems, Comment{Text: strings.Trim(comment[0][1], " ")})
	}
	return lineItems
}
