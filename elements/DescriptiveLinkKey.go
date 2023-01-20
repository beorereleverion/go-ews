package elements

// The DescriptiveLinkKey element is currently unused and is reserved for future use. It contains a value of 0.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/descriptivelinkkey
import "encoding/xml"

type DescriptiveLinkKey struct {
	XMLName xml.Name
	TEXT    int64 `xml:",chardata"`
}

func (D *DescriptiveLinkKey) SetForMarshal() {
	D.XMLName.Local = "m:DescriptiveLinkKey"
}

func (D *DescriptiveLinkKey) GetSchema() *Schema {
	return &SchemaMessages
}
