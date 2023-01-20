package elements

// The InstanceKey element specifies an instance key for an item or conversation.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/instancekey
import "encoding/xml"

type InstanceKey struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (I *InstanceKey) SetForMarshal() {
	I.XMLName.Local = "t:InstanceKey"
}

func (I *InstanceKey) GetSchema() *Schema {
	return &SchemaTypes
}
