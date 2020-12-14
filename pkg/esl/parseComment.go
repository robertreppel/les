package esl

import (
	"regexp"
	"strings"
)

func parseComment(eslInput string, lineItems []Item) []Item {
	re, _ := regexp.Compile("^ *// *(.*)")
	comment := re.FindAllStringSubmatch(eslInput, -1)
	if len(comment) > 0 {
		commentText := strings.Trim(comment[0][1], " ")
		commentText = strings.TrimLeft(commentText, "/")
		commentText = strings.TrimLeft(commentText, " ")
		lineItems = append(lineItems, Comment{Text: commentText, Type: "Comment"})
	}
	return lineItems
}
