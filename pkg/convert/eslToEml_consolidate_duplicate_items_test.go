package convert_test

import (
	"fmt"
	"testing"

	"github.com/robertreppel/les/pkg/convert"
	"github.com/robertreppel/les/pkg/esl"
)

func TestNoDuplicateEventsShouldLeadToSingleEvent(t *testing.T) {
	input := []string{
		"User Registered",
	}
	markdown, err := esl.Parse(input)
	if err != nil {
		panic(err)
	}
	markup, err := convert.EslToEml(markdown)
	if err != nil {
		panic(err)
	}

	streams := markup.Esl.Contexts[0].Streams
	fmt.Println(streams[0].Events)
	if len(streams[0].Events) != 1 {
		t.Error("expected different no of User events.")
	}
}
