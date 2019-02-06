package main

import (
	"fmt"
	"os"

	"gopkg.in/alecthomas/kingpin.v2"
)

const defaultEslFile = "Eventstorming.esl"

func main() {
	app := kingpin.New("les-viz", "Generates a http://www.graphviz.org/ digraph for an event storming.")
	app.Version("0.10.7-alpha")
	file := app.Arg("file", "Event Storming Language (.esl) file").String()
	kingpin.MustParse(app.Parse(os.Args[1:]))
	eslToGraphVizDigraph(*file)
}

func printError(id string, message string) {
	fmt.Printf("%s: %s\n", id, message)
}
