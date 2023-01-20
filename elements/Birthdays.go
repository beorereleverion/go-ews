package elements

// The Birthdays element specifies an array of birthdays, stored as strings, and the identifiers of their source attributions for the associated persona.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/birthdays
import "encoding/xml"

type Birthdays struct {
	XMLName xml.Name

	// The StringAttributedValue element specifies an instance in an array of attributes associated with a persona element.
	StringAttributedValue *StringAttributedValue `xml:"StringAttributedValue"`
}

func (B *Birthdays) SetForMarshal() {
	B.XMLName.Local = "t:Birthdays"
}

func (B *Birthdays) GetSchema() *Schema {
	return &SchemaTypes
}
