package elements

// The OfficeLocation element represents the office location of a contact.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/officelocation
import "encoding/xml"

type OfficeLocation struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (O *OfficeLocation) SetForMarshal() {
	O.XMLName.Local = "t:OfficeLocation"
}

func (O *OfficeLocation) GetSchema() *Schema {
	return &SchemaTypes
}
