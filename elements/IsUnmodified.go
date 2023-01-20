package elements

// The IsUnmodified element indicates whether the item has been modified.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/isunmodified
import "encoding/xml"

type IsUnmodified struct {
	XMLName xml.Name
	TEXT    bool `xml:",chardata"`
}

func (I *IsUnmodified) SetForMarshal() {
	I.XMLName.Local = "t:IsUnmodified"
}

func (I *IsUnmodified) GetSchema() *Schema {
	return &SchemaTypes
}
