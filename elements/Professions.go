package elements

// The Professions element specifies an array of Profession values and the identifiers of their source attributions for the associated persona.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/professions
import "encoding/xml"

type Professions struct {
	XMLName xml.Name

	// The StringAttributedValue element specifies an instance in an array of attributes associated with a persona element.
	StringAttributedValue *StringAttributedValue `xml:"StringAttributedValue"`
}

func (P *Professions) SetForMarshal() {
	P.XMLName.Local = "t:Professions"
}

func (P *Professions) GetSchema() *Schema {
	return &SchemaTypes
}
