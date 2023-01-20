package elements

// The Suffix element represents a suffix to a contact's name.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/suffix
import "encoding/xml"

type Suffix struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (S *Suffix) SetForMarshal() {
	S.XMLName.Local = "t:Suffix"
}

func (S *Suffix) GetSchema() *Schema {
	return &SchemaTypes
}
