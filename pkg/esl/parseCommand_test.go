package esl_test

import (
	"testing"

	"github.com/robertreppel/les/pkg/esl"
)

func TestShouldGetItemOfTypeCommand(t *testing.T) {
	input := []string{"  Validate Registration  -> : userId                     "}
	result, err := esl.Parse(input)
	if err != nil {
		panic(err)
	}
	if len(result.Lines) == 0 {
		t.Error("no command found")
		return
	}
	switch result.Lines[0].(type) {
	case esl.Command:
		if result.Lines[0].(esl.Command).Type != "Command" {
			t.Error("Unexpected Command.Type")
		}
	default:
		t.Error("expected command")
	}
}

func TestShouldGetCommandWithParameters(t *testing.T) {
	input := []string{"  Validate Registration  -> : userId                     "}
	result, err := esl.Parse(input)
	if err != nil {
		panic(err)
	}
	if len(result.Lines) == 0 {
		t.Error("no command found")
		return
	}
	switch result.Lines[0].(type) {
	case esl.Command:
		if result.Lines[0].(esl.Command).Name != "Validate Registration" {
			t.Error("Unexpected Command.Name")
		}
	default:
		t.Error("expected command")
	}
}

func Test_should_get_command_with_parameters_without_slashes(t *testing.T) {
	input := []string{"  Validate Registration  -> userId                     "}
	result, err := esl.Parse(input)
	if err != nil {
		panic(err)
	}
	if len(result.Lines) == 0 {
		t.Error("no command found")
		return
	}
	switch result.Lines[0].(type) {
	case esl.Command:
		if result.Lines[0].(esl.Command).Name != "Validate Registration" {
			t.Log(result)
			t.Error("Unexpected Command.Name")
		}
	default:
		t.Error("expected command")
	}
}

func TestShouldFindCommandParameters(t *testing.T) {
	input := []string{"Validate->: userId  ,  "}
	result, err := esl.Parse(input)
	if err != nil {
		panic(err)
	}
	if len(result.Errors) > 0 {
		t.Log(result.Errors)
		t.Error("Unexpected error(s).")
	}
	if len(result.Lines) == 0 {
		t.Error("no command found")
		return
	}
	switch result.Lines[0].(type) {
	case esl.Command:
		parameters := result.Lines[0].(esl.Command).Parameters
		if len(parameters) != 1 {
			t.Log(parameters)
			t.Error("Unexpected number of command.Parameters")
		}
		if parameters[0].Name != "userId" {
			t.Error("Command parameter not found")
		}
	default:
		t.Error("expected command")
	}
}

func TestShouldGetCommandWithoutProperties(t *testing.T) {
	input := []string{"Register->                   "}
	result, err := esl.Parse(input)
	if err != nil {
		panic(err)
	}
	if len(result.Lines) == 0 {
		t.Error("no event found")
		return
	}
	switch result.Lines[0].(type) {
	case esl.Command:
		if result.Lines[0].(esl.Command).Name != "Register" {
			t.Error("Unexpected Command.Name")
		}
	default:
		t.Error("expected command")
	}
}

func TestShouldGetCommandWithoutPropertiesWithTrailingSlashes(t *testing.T) {
	input := []string{"Register->:  "}
	result, err := esl.Parse(input)
	if err != nil {
		panic(err)
	}
	if len(result.Lines) == 0 {
		t.Error("no command found")
		return
	}
	switch result.Lines[0].(type) {
	case esl.Command:
		if result.Lines[0].(esl.Command).Name != "Register" {
			t.Error("Unexpected Command.Name")
		}
	default:
		t.Error("expected command")
	}
}
