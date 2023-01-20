package elements

// The Attribution element specifies a string used to identify an attribute of a persona.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/attribution-string
import "encoding/xml"

type Attributionstring struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (A *Attributionstring) SetForMarshal() {
	A.XMLName.Local = "t:Attribution"
}

func (A *Attributionstring) GetSchema() *Schema {
	return &SchemaTypes
}
