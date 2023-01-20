package elements

// The PostalCode element represents the postal code for a contact item.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/postalcode
import "encoding/xml"

type PostalCode struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (P *PostalCode) SetForMarshal() {
	P.XMLName.Local = "t:PostalCode"
}

func (P *PostalCode) GetSchema() *Schema {
	return &SchemaTypes
}
