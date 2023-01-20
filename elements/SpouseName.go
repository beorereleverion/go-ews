package elements

// The SpouseName element represents the name of a contact's spouse or partner.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/spousename
import "encoding/xml"

type SpouseName struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (S *SpouseName) SetForMarshal() {
	S.XMLName.Local = "t:SpouseName"
}

func (S *SpouseName) GetSchema() *Schema {
	return &SchemaTypes
}
