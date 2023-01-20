package elements

// The People element specifies an array of persona data returned as the result of a FindPeople request.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/people
import "encoding/xml"

type People struct {
	XMLName xml.Name

	// The Persona element specifies a set of persona data returned by a GetPersona request.
	Persona *Persona `xml:"Persona"`
}

func (P *People) SetForMarshal() {
	P.XMLName.Local = "m:People"
}

func (P *People) GetSchema() *Schema {
	return &SchemaMessages
}
