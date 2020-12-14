package esl_test

import (
	"testing"

	"github.com/robertreppel/les/pkg/esl"
)

func TestShouldGetItemOfTypeComment(t *testing.T) {
	input := []string{"// commenting at length "}
	result, err := esl.Parse(input)
	if err != nil {
		panic(err)
	}
	if len(result.Lines) == 0 {
		t.Error("no comment found")
		return
	}
	switch result.Lines[0].(type) {
	case esl.Comment:
		if result.Lines[0].(esl.Comment).Type != "Comment" {
			t.Error("Unexpected Comment.Type")
		}
	default:
		t.Error("expected command")
	}
}

func TestComments(t *testing.T) {
	input := []string{"// This is a comment                     "}
	result, err := esl.Parse(input)
	if err != nil {
		panic(err)
	}

	switch result.Lines[0].(type) {
	case esl.Comment:
		if result.Lines[0].(esl.Comment).Text != "This is a comment" {
			t.Error("Unexpected Comment.Text")
		}
	default:
		t.Error("expected comment")

	}
}
