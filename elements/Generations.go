package elements

// The Generations element specifies an array of generation values and the identifiers of their source attributions for the associated persona.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/generations
import "encoding/xml"

type Generations struct {
	XMLName xml.Name

	// The StringAttributedValue element specifies an instance in an array of attributes associated with a persona element.
	StringAttributedValue *StringAttributedValue `xml:"StringAttributedValue"`
}

func (G *Generations) SetForMarshal() {
	G.XMLName.Local = "t:Generations"
}

func (G *Generations) GetSchema() *Schema {
	return &SchemaTypes
}
