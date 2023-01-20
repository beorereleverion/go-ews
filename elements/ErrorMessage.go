package elements

// The ErrorMessage element represents the reason for the validation error.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/errormessage
import "encoding/xml"

type ErrorMessage struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (E *ErrorMessage) SetForMarshal() {
	E.XMLName.Local = "m:ErrorMessage"
}

func (E *ErrorMessage) GetSchema() *Schema {
	return &SchemaMessages
}
