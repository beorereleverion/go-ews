package elements

// The MessageText element provides a text description of the status of the response.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/messagetext
import "encoding/xml"

type MessageText struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (M *MessageText) SetForMarshal() {
	M.XMLName.Local = "m:MessageText"
}

func (M *MessageText) GetSchema() *Schema {
	return &SchemaMessages
}
