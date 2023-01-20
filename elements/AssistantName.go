package elements

// The AssitantName element represents an assistant to a contact.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/assistantname
import "encoding/xml"

type AssistantName struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (A *AssistantName) SetForMarshal() {
	A.XMLName.Local = "t:AssistantName"
}

func (A *AssistantName) GetSchema() *Schema {
	return &SchemaTypes
}
