package convert_test

import (
	"fmt"
	"testing"

	"github.com/robertreppel/les/pkg/convert"
	"github.com/robertreppel/les/pkg/esl"
)

func TestDuplicateEventsShouldBeConsolidated(t *testing.T) {
	input := []string{
		"User Registered",
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
	if len(streams[0].Events) != 1 {
		t.Error("expected different no of User events.")
	}
}

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

func TestDuplicateEventPropertiesShouldBeUnioned(t *testing.T) {
	input := []string{
		"User Registered",
		"User Registered : email, userId",
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
	if len(streams[0].Events) != 1 {
		t.Error("expected different no of User events.")
	}

	event := streams[0].Events[0]
	if len(event.Event.Properties) != 2 {
		t.Error("Expected different number of properties.")
	}

}
