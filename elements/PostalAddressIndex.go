package elements

// The PostalAddressIndex element represents the display types for physical addresses.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/postaladdressindex
import "encoding/xml"

type PostalAddressIndex struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (P *PostalAddressIndex) SetForMarshal() {
	P.XMLName.Local = "t:PostalAddressIndex"
}

func (P *PostalAddressIndex) GetSchema() *Schema {
	return &SchemaTypes
}
