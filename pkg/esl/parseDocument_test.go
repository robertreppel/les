package esl_test

import (
	"testing"

	"github.com/robertreppel/les/pkg/esl"
)

func TestShouldGetDocumentWithParameters(t *testing.T) {
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

func TestShouldGetDocumentParameters(t *testing.T) {
	input := []string{"Users* : userId,email"}
	result, err := esl.Parse(input)
	if err != nil {
		panic(err)
	}
	if len(result.Lines) == 0 {
		t.Error("no document found")
		return
	}
	switch result.Lines[0].(type) {
	case esl.Document:
		properties := result.Lines[0].(esl.Document).Properties
		if len(properties) != 2 {
			t.Error("Unexpected number of document.Properties")
		}
		if properties[0].Name != "userId" {
			t.Error("Document property not found")
		}
		if properties[1].Name != "email" {
			t.Error("Document property not found")
		}
	default:
		t.Error("expected document")
	}
}

func TestShouldGetDocumentParametersWithTrailingComma(t *testing.T) {
	input := []string{"Users* : userId,email  ,     "}
	result, err := esl.Parse(input)
	if err != nil {
		panic(err)
	}
	if len(result.Lines) == 0 {
		t.Error("no document found")
		return
	}
	switch result.Lines[0].(type) {
	case esl.Document:
		properties := result.Lines[0].(esl.Document).Properties
		if len(properties) != 2 {
			t.Error("Unexpected number of document.Properties")
		}
		if properties[0].Name != "userId" {
			t.Error("Document property not found")
		}
		if properties[1].Name != "email" {
			t.Error("Document property not found")
		}
	default:
		t.Error("expected document")
	}
}

func TestShouldGetDocumentProperties(t *testing.T) {
	input := []string{"Register* : userId,email,password                     "}
	result, err := esl.Parse(input)
	if err != nil {
		panic(err)
	}
	if len(result.Lines) == 0 {
		t.Error("no document found")
		return
	}
	switch result.Lines[0].(type) {
	case esl.Document:
		properties := result.Lines[0].(esl.Document).Properties
		if len(properties) != 3 {
			t.Error("Unexpected number of Document.Properties")
		}
		if properties[0].Name != "userId" {
			t.Error("Document property not found")
		}
		if properties[1].Name != "email" {
			t.Error("Document property not found")
		}
		if properties[2].Name != "password" {
			t.Error("Document property not found")
		}
	default:
		t.Error("expected document")
	}
}

func TestShouldGetDocumentWithoutProperties(t *testing.T) {
	input := []string{"User Register*                    "}
	result, err := esl.Parse(input)
	if err != nil {
		panic(err)
	}
	if len(result.Lines) == 0 {
		t.Error("no document found")
		return
	}
	switch result.Lines[0].(type) {
	case esl.Document:
		if result.Lines[0].(esl.Document).Name != "User Register" {
			t.Error("Unexpected Document.Name")
		}
	default:
		t.Error("expected document")
	}
}

func TestShouldGetDocumentWithoutPropertiesWithTrailingSlashes(t *testing.T) {
	input := []string{"A List* :"}
	result, err := esl.Parse(input)
	if err != nil {
		panic(err)
	}
	if len(result.Lines) == 0 {
		t.Error("no document found")
		return
	}
	switch result.Lines[0].(type) {
	case esl.Document:
		if result.Lines[0].(esl.Document).Name != "A List" {
			t.Error("Unexpected Document.Name")
		}
	default:
		t.Error("expected document")
	}
}
