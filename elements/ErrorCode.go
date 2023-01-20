package elements

// The ErrorCode element represents a rule validation error code that describes what failed validation for each rule predicate or action.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/errorcode
import "encoding/xml"

type ErrorCode struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (E *ErrorCode) SetForMarshal() {
	E.XMLName.Local = "m:ErrorCode"
}

func (E *ErrorCode) GetSchema() *Schema {
	return &SchemaMessages
}
