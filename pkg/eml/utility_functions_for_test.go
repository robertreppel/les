package eml_test

import (
	"github.com/robertreppel/les/pkg/eml"
)

func hasError(errorId string, errors []eml.ValidationError) bool {
	for _, err := range errors {
		if err.ErrorID == errorId {
			return true
		}
	}
	return false
}
