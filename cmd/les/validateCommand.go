package main

import (
	"fmt"
	"os"

	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

type validateCommand struct {
	tail      *bool
	inputType *string
	file      *string
}

func configureValidateCommand(app *kingpin.Application) {
	c := &validateCommand{}
	cmd := app.Command("validate", "Verify if event markdown in ./"+defaultEslFile+" or event markup in ./"+defaultEmlFile+" is valid for building APIs.")
	c.tail = cmd.Flag("follow", "Re-validate whenever ./"+defaultEslFile+" or  ./"+defaultEmlFile+" are changed.").Short('f').Bool()
	c.file = cmd.Arg("file", ".esl or .eml.yaml file. Default: "+defaultEslFile+" or "+defaultEmlFile+".").String()
	cmd.Action(c.validate)
}

func (n *validateCommand) validate(c *kingpin.ParseContext) error {
	inputFile := useDefaultEslOrEmlFileIfInputFileNotSpecified(*n.file)
	if inputFile == "" {
		fmt.Println("No input file found. Try 'les validate --help'.")
		return nil
	}
	tail := false
	if n.tail != nil {
		tail = *n.tail
	}
	if tail {
		whenFileChangesThenValidate(inputFile, ".")
		return nil
	}

	if !isFileValidEslOrEml(inputFile) {
		os.Exit(-1)
	}
	return nil
}
