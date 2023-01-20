package elements

// The GetPersona element contains the request to get a persona.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/getpersona
import "encoding/xml"

type GetPersona struct {
	XMLName xml.Name

	// The PersonaId element specifies the persona identifier for the associated persona.
	PersonaId *PersonaId `xml:"PersonaId"`
}

func (G *GetPersona) SetForMarshal() {
	G.XMLName.Local = "m:GetPersona"
}

func (G *GetPersona) GetSchema() *Schema {
	return &SchemaMessages
}
