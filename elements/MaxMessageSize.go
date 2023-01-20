package elements

// The MaxMessageSize element represents the maximum message size a recipient can accept.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/maxmessagesize
import "encoding/xml"

type MaxMessageSize struct {
	XMLName xml.Name
	TEXT    int64 `xml:",chardata"`
}

func (M *MaxMessageSize) SetForMarshal() {
	M.XMLName.Local = "t:MaxMessageSize"
}

func (M *MaxMessageSize) GetSchema() *Schema {
	return &SchemaTypes
}
