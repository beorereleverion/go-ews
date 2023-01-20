package elements

// The Hobbies element specifies an array of hobbies and the identifiers of their source attributions for the associated persona.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/hobbies
import "encoding/xml"

type Hobbies struct {
	XMLName xml.Name

	// The StringAttributedValue element specifies an instance in an array of attributes associated with a persona element.
	StringAttributedValue *StringAttributedValue `xml:"StringAttributedValue"`
}

func (H *Hobbies) SetForMarshal() {
	H.XMLName.Local = "t:Hobbies"
}

func (H *Hobbies) GetSchema() *Schema {
	return &SchemaTypes
}
