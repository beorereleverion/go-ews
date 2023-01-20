package elements

// The PersonaId element specifies the persona identifier for the associated persona.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/personaid
import "encoding/xml"

type PersonaId struct {
	XMLName xml.Name

	// The text value of the ChangeKey attribute is the change key of the persona.
	ChangeKey *string `xml:"ChangeKey,attr"`
	// The text value of the Id attribute is the identifier of the persona.
	Id *string `xml:"Id,attr"`
}

func (P *PersonaId) SetForMarshal() {
	P.XMLName.Local = "m:PersonaId"
}

func (P *PersonaId) GetSchema() *Schema {
	return &SchemaMessages
}
