package elements

// The MessageXml element provides additional error response information.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/messagexml
import "encoding/xml"

type MessageXml struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (M *MessageXml) SetForMarshal() {
	M.XMLName.Local = "m:MessageXml"
}

func (M *MessageXml) GetSchema() *Schema {
	return &SchemaMessages
}
