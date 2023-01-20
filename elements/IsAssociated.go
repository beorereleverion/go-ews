package elements

// The IsAssociated element indicates whether the item is associated with a folder.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/isassociated
import "encoding/xml"

type IsAssociated struct {
	XMLName xml.Name
	TEXT    bool `xml:",chardata"`
}

func (I *IsAssociated) SetForMarshal() {
	I.XMLName.Local = "t:IsAssociated"
}

func (I *IsAssociated) GetSchema() *Schema {
	return &SchemaTypes
}
