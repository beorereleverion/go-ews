package elements

// The Department element represents the contact's department at work.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/department
import "encoding/xml"

type Department struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (D *Department) SetForMarshal() {
	D.XMLName.Local = "t:Department"
}

func (D *Department) GetSchema() *Schema {
	return &SchemaTypes
}
