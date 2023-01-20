package elements

// The PersonName element specifies the name of an individual found by means of entity extraction.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/personname
import "encoding/xml"

type PersonName struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (P *PersonName) SetForMarshal() {
	P.XMLName.Local = "t:PersonName"
}

func (P *PersonName) GetSchema() *Schema {
	return &SchemaTypes
}
