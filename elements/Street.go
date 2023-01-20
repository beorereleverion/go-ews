package elements

// The Street element represents a street address for a contact item.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/street
import "encoding/xml"

type Street struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (S *Street) SetForMarshal() {
	S.XMLName.Local = "t:Street"
}

func (S *Street) GetSchema() *Schema {
	return &SchemaTypes
}
