package elements

// The FirstName element represents the first name of a contact.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/firstname
import "encoding/xml"

type FirstName struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (F *FirstName) SetForMarshal() {
	F.XMLName.Local = "t:FirstName"
}

func (F *FirstName) GetSchema() *Schema {
	return &SchemaTypes
}
