package elements

// The Error element represents a single validation error on a particular rule property value, predicate property value, or action property value.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/error
import "encoding/xml"

type Error struct {
	XMLName xml.Name

	// The ErrorCode element represents a rule validation error code that describes what failed validation for each rule predicate or action.
	ErrorCode *ErrorCode `xml:"ErrorCode"`
	// The ErrorMessage element represents the reason for the validation error.
	ErrorMessage *ErrorMessage `xml:"ErrorMessage"`
	// The FieldURI element specifies the URI to the rule field that caused the validation error.
	FieldUri *FieldUriRule `xml:"FieldUri"`
	// The FieldValue element represents the value of the field that caused the validation error.
	FieldValue *FieldValue `xml:"FieldValue"`
}

func (E *Error) SetForMarshal() {
	E.XMLName.Local = "m:Error"
}

func (E *Error) GetSchema() *Schema {
	return &SchemaMessages
}
