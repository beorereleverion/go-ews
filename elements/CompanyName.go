package elements

// The CompanyName element represents the company name that is associated with a contact.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/companyname
import "encoding/xml"

type CompanyName struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (C *CompanyName) SetForMarshal() {
	C.XMLName.Local = "t:CompanyName"
}

func (C *CompanyName) GetSchema() *Schema {
	return &SchemaTypes
}
