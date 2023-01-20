package elements

// The RequestedExtensionIds element contains an array of extension identifiers.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/requestedextensionids
import "encoding/xml"

type RequestedExtensionIds struct {
	XMLName xml.Name

	// The String element represents a string that is used by items, contacts, tasks, and conversations.
	String *String `xml:"String"`
}

func (R *RequestedExtensionIds) SetForMarshal() {
	R.XMLName.Local = "m:RequestedExtensionIds"
}

func (R *RequestedExtensionIds) GetSchema() *Schema {
	return &SchemaMessages
}
