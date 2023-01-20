package elements

// The Name element represents the property name for a message tracking report.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/name-message-tracking
import "encoding/xml"

type NameMessageTracking struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (N *NameMessageTracking) SetForMarshal() {
	N.XMLName.Local = "t:Name"
}

func (N *NameMessageTracking) GetSchema() *Schema {
	return &SchemaTypes
}
