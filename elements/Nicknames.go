package elements

// The Nicknames element specifies an array of nickname values and the identifiers of their source attributions for the associated persona.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/nicknames
import "encoding/xml"

type Nicknames struct {
	XMLName xml.Name

	// The StringAttributedValue element specifies an instance in an array of attributes associated with a persona element.
	StringAttributedValue *StringAttributedValue `xml:"StringAttributedValue"`
}

func (N *Nicknames) SetForMarshal() {
	N.XMLName.Local = "t:Nicknames"
}

func (N *Nicknames) GetSchema() *Schema {
	return &SchemaTypes
}
