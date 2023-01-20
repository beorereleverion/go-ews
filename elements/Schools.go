package elements

// The Schools element specifies an array of school names and the identifiers of their source attributions for the associated persona.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/schools
import "encoding/xml"

type Schools struct {
	XMLName xml.Name

	// The StringAttributedValue element specifies an instance in an array of attributes associated with a persona element.
	StringAttributedValue *StringAttributedValue `xml:"StringAttributedValue"`
}

func (S *Schools) SetForMarshal() {
	S.XMLName.Local = "t:Schools"
}

func (S *Schools) GetSchema() *Schema {
	return &SchemaTypes
}
