package elements

// The Attributions element specifies an array of attributions for its associated Value element.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/attributions-arrayofvalueattributionstype
import "encoding/xml"

type AttributionsArrayOfValueAttributionsType struct {
	XMLName xml.Name

	// The Attribution element specifies a string used to identify an attribute of a persona.
	Attribution *Attributionstring `xml:"Attribution"`
}

func (A *AttributionsArrayOfValueAttributionsType) SetForMarshal() {
	A.XMLName.Local = "t:Attributions"
}

func (A *AttributionsArrayOfValueAttributionsType) GetSchema() *Schema {
	return &SchemaTypes
}
