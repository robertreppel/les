# les-viz: Event Storming Language Vizualization (ESL -> GraphViz)

![example](consentaur-example.png)

**Generates a [GraphViz](http://www.graphviz.org/) digraph for an event storming.**

[What is Event Storming Language?](https://docs.letseventsource.org/faq/eventmarkdown/)

See also: https://webeventstorming.com

## Usage

```bash

$ les-viz --help
usage: les-viz [<flags>] [<file>]

Generates a http://www.graphviz.org/ digraph for an event storming.

Flags:
  --help     Show context-sensitive help (also try --help-long and --help-man).
  --version  Show application version.

Args:
  [<file>]  Event Storming Language (.esl) file

```

## Generating a .png Image

Install GraphViz:

```bash

sudo apt-get install graphviz

```

Generate a visual representation of an event storming:

```bash

curl -L https://raw.githubusercontent.com/Adaptech/les/master/samples/consentaur/Eventstorming.esl > Eventstorming.esl \
&& les-viz Eventstorming.esl | dot -Tpng -o Eventstorming.png \
&& eog Eventstorming.png

```
