package elements

// The Value element represents the property value for a message tracking report.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/value-message-tracking
import "encoding/xml"

type ValueMessageTracking struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (V *ValueMessageTracking) SetForMarshal() {
	V.XMLName.Local = "t:Value"
}

func (V *ValueMessageTracking) GetSchema() *Schema {
	return &SchemaTypes
}
