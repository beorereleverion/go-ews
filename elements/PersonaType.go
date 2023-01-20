package elements

// The PersonaType element specifies the type of the persona, for example, a person or a distribution list.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/personatype
import "encoding/xml"

type PersonaType struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (P *PersonaType) SetForMarshal() {
	P.XMLName.Local = "t:PersonaType"
}

func (P *PersonaType) GetSchema() *Schema {
	return &SchemaTypes
}
