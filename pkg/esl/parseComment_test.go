package esl_test

import (
	"testing"

	"github.com/robertreppel/les/pkg/esl"
)

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
