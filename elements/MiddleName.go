package elements

// The MiddleName element represents the middle name of a contact.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/middlename
import "encoding/xml"

type MiddleName struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (M *MiddleName) SetForMarshal() {
	M.XMLName.Local = "t:MiddleName"
}

func (M *MiddleName) GetSchema() *Schema {
	return &SchemaTypes
}
