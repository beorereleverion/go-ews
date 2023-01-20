package elements

// The IsFromMe element indicates whether a user sent an item to him or herself.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/isfromme
import "encoding/xml"

type IsFromMe struct {
	XMLName xml.Name
	TEXT    bool `xml:",chardata"`
}

func (I *IsFromMe) SetForMarshal() {
	I.XMLName.Local = "t:IsFromMe"
}

func (I *IsFromMe) GetSchema() *Schema {
	return &SchemaTypes
}
