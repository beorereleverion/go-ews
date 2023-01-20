package elements

// The PostOfficeBox element specifies thepost office boxportion of a postal address.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/postofficebox
import "encoding/xml"

type PostOfficeBox struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (P *PostOfficeBox) SetForMarshal() {
	P.XMLName.Local = "t:PostOfficeBox"
}

func (P *PostOfficeBox) GetSchema() *Schema {
	return &SchemaTypes
}
