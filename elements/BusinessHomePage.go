package elements

// The BusinessHomePage element represents the Home page (Web address) for the contact.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/businesshomepage
import "encoding/xml"

type BusinessHomePage struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (B *BusinessHomePage) SetForMarshal() {
	B.XMLName.Local = "t:BusinessHomePage"
}

func (B *BusinessHomePage) GetSchema() *Schema {
	return &SchemaTypes
}
