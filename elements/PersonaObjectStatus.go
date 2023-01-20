package elements

// The PersonaObjectStatus element specifies whether the information in the associated persona is complete or partial.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/personaobjectstatus
import "encoding/xml"

type PersonaObjectStatus struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (P *PersonaObjectStatus) SetForMarshal() {
	P.XMLName.Local = "t:PersonaObjectStatus"
}

func (P *PersonaObjectStatus) GetSchema() *Schema {
	return &SchemaTypes
}
