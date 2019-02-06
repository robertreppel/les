package esl

// Esl is a specification written in Event Storming Language.
type Esl struct {
	Lines  []Item
	Errors []EslValidationError `yaml:"Errors"` // The presence of errors means that no conversion to

}

// Item is a line item in an Event Storming Language file.
type Item interface {
}

// A Comment in Event Storming Language
// Example:
// "Order Placed"
type Comment struct {
	Text string
}

// Command in Event Storming Language language
// Example:
// "Place Order-> // orderId, placedDate, deliveryDate"
type Command struct {
	Name       string
	Parameters []Parameter
}

// A Parameter for an esl command
type Parameter struct {
	Name string
}

// Event describes an esl event
type Event struct {
	Name       string
	Properties []Property
}

// Document describes an esl read model document
type Document struct {
	Name       string
	Properties []Property
}

// A Property of an esl event
type Property struct {
	Name string
}

// A EslValidationError means that the event markdown structure cannot be used to generate event markup.
type EslValidationError struct {
	ErrorID string
	Message string
}
