package elements

// The Attributions element specifies an array of attribution information for one or more of the contacts or Active Directory recipients aggregated into the associated persona.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/attributions-arrayofpersonaattributionstype
import "encoding/xml"

type AttributionsArrayOfPersonaAttributionsType struct {
	XMLName xml.Name

	// The Attribution element specifies an instance in an array of attributes for a PersonaType element.
	Attribution *AttributionPersonaAttributionType `xml:"Attribution"`
}

func (A *AttributionsArrayOfPersonaAttributionsType) SetForMarshal() {
	A.XMLName.Local = "t:Attributions"
}

func (A *AttributionsArrayOfPersonaAttributionsType) GetSchema() *Schema {
	return &SchemaTypes
}
