package main

import (
	"fmt"
	"log"

	"github.com/robertreppel/les/pkg/convert"
	"github.com/robertreppel/les/pkg/eml"
	"github.com/robertreppel/les/pkg/esl"
)

func convertFileToEml(inputFile string, outputFile string) (bool, error) {
	input, err := ReadLines(inputFile)
	if err != nil {
		return false, fmt.Errorf("convertToEml: %v", err)
	}

	markdown, err := esl.Parse(input)
	if err != nil {
		return false, fmt.Errorf("convertToEml: %v", err)
	}
	conversionResult, err := convert.EslToEml(markdown)
	if err != nil {
		return false, fmt.Errorf("convertToEml: %v", err)
	}

	isValidEml := len(conversionResult.MarkdownValidationErrors) == 0
	if !isValidEml {
		fmt.Println("ESL Errors:")
		for _, eslErr := range conversionResult.MarkdownValidationErrors {
			printError(eslErr.ErrorID, eslErr.Message)
		}
	}

	markup := conversionResult.Esl
	yaml, err := eml.ToYaml(markup)
	if err != nil {
		log.Panicf("convertFileToEml: %v", err)
	}
	err = WriteToFile(outputFile, yaml)
	if err != nil {
		log.Panicf("convertFileToEml: %v", err)
	}
	return isValidEml, nil
}
