package esl_test

import (
	"testing"

	"github.com/robertreppel/les/pkg/esl"
)

func TestShouldGetEventWithProperties(t *testing.T) {
	input := []string{"User Registered // userId, email, password                     "}
	result, err := esl.Parse(input)
	if err != nil {
		panic(err)
	}
	if len(result.Lines) == 0 {
		t.Error("no event found")
		return
	}
	switch result.Lines[0].(type) {
	case esl.Event:
		if result.Lines[0].(esl.Event).Name != "User Registered" {
			t.Error("Unexpected Event.Name")
		}
	default:
		t.Error("expected event")
	}
}

func TestShouldGetProperties(t *testing.T) {
	input := []string{"User Registered // userId,email,password                     "}
	result, err := esl.Parse(input)
	if err != nil {
		panic(err)
	}
	if len(result.Lines) == 0 {
		t.Error("no event found")
		return
	}
	switch result.Lines[0].(type) {
	case esl.Event:
		properties := result.Lines[0].(esl.Event).Properties
		if len(properties) != 3 {
			t.Error("Unexpected number of Event.Properties")
		}
		if properties[0].Name != "userId" {
			t.Error("Event property not found")
		}
		if properties[1].Name != "email" {
			t.Error("Event property not found")
		}
		if properties[2].Name != "password" {
			t.Error("Event property not found")
		}
	default:
		t.Error("expected event")
	}
}

func TestShouldGetPropertiesWithTrailingComma(t *testing.T) {
	input := []string{"User Registered // userId,email,password,                     "}
	result, err := esl.Parse(input)
	if err != nil {
		panic(err)
	}
	if len(result.Lines) == 0 {
		t.Error("no event found")
		return
	}
	switch result.Lines[0].(type) {
	case esl.Event:
		properties := result.Lines[0].(esl.Event).Properties
		if len(properties) != 3 {
			t.Error("Unexpected number of Event.Properties")
		}
		if properties[0].Name != "userId" {
			t.Error("Event property not found")
		}
		if properties[1].Name != "email" {
			t.Error("Event property not found")
		}
		if properties[2].Name != "password" {
			t.Error("Event property not found")
		}
	default:
		t.Error("expected event")
	}
}

func TestShouldGetEventWithoutProperties(t *testing.T) {
	input := []string{"User Registered                    "}
	result, err := esl.Parse(input)
	if err != nil {
		panic(err)
	}
	if len(result.Lines) == 0 {
		t.Error("no event found")
		return
	}
	switch result.Lines[0].(type) {
	case esl.Event:
		if result.Lines[0].(esl.Event).Name != "User Registered" {
			t.Error("Unexpected Event.Name")
		}
	default:
		t.Error("expected event")
	}
}

func TestShouldGetEventWithoutPropertiesWithTrailingSlashes(t *testing.T) {
	input := []string{"User Registered //  "}
	result, err := esl.Parse(input)
	if err != nil {
		panic(err)
	}
	if len(result.Lines) == 0 {
		t.Error("no event found")
		return
	}
	switch result.Lines[0].(type) {
	case esl.Event:
		if result.Lines[0].(esl.Event).Name != "User Registered" {
			t.Error("Unexpected Event.Name")
		}
		if len(result.Lines[0].(esl.Event).Properties) != 0 {
			t.Error("Unexpected Event.Properties.")
		}
	default:
		t.Error("expected event")
	}
}

func TestShouldNotReturnEventWhenCommandGiven(t *testing.T) {
	input := []string{"Command This->"}
	result, err := esl.Parse(input)
	if err != nil {
		panic(err)
	}
	if len(result.Lines) > 0 {
		switch result.Lines[0].(type) {
		case esl.Event:
			t.Error("Unexpected Event.")
		}
	}
}
func TestShouldNotReturnEventWhenDocumentGiven(t *testing.T) {
	input := []string{"A Read Model* // one,two"}
	result, err := esl.Parse(input)
	if err != nil {
		panic(err)
	}
	if len(result.Lines) > 0 {
		switch result.Lines[0].(type) {
		case esl.Event:
			t.Error("Unexpected Event.")
		}
	}
}
