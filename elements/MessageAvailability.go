package elements

// The Message element contains the out of Office (OOF) response.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/message-availability
import "encoding/xml"

type MessageAvailability struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (M *MessageAvailability) SetForMarshal() {
	M.XMLName.Local = "t:Message"
}

func (M *MessageAvailability) GetSchema() *Schema {
	return &SchemaTypes
}
