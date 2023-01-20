package elements

// The IsInError element indicates whether the rule is in an error condition.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/isinerror
import "encoding/xml"

type IsInError struct {
	XMLName xml.Name
	TEXT    bool `xml:",chardata"`
}

func (I *IsInError) SetForMarshal() {
	I.XMLName.Local = "m:IsInError"
}

func (I *IsInError) GetSchema() *Schema {
	return &SchemaMessages
}
