package elements

// The MessageId element represents the message identification to search for.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/messageid
import "encoding/xml"

type MessageId struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (M *MessageId) SetForMarshal() {
	M.XMLName.Local = "m:MessageId"
}

func (M *MessageId) GetSchema() *Schema {
	return &SchemaMessages
}
