package elements

// The ValidationErrors element represents an array of rule validation errors on each rule field that has an error.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/validationerrors
import "encoding/xml"

type ValidationErrors struct {
	XMLName xml.Name

	// The Error element represents a single validation error on a particular rule property value, predicate property value, or action property value.
	Error *Error `xml:"Error"`
}

func (V *ValidationErrors) SetForMarshal() {
	V.XMLName.Local = "m:ValidationErrors"
}

func (V *ValidationErrors) GetSchema() *Schema {
	return &SchemaMessages
}
