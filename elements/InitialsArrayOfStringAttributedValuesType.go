package elements

// The Initials element specifies an array of initials values and the identifiers of their source attributions for the associated persona.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/initials-arrayofstringattributedvaluestype
import "encoding/xml"

type InitialsArrayOfStringAttributedValuesType struct {
	XMLName xml.Name

	// The StringAttributedValue element specifies an instance in an array of attributes associated with a persona element.
	StringAttributedValue *StringAttributedValue `xml:"StringAttributedValue"`
}

func (I *InitialsArrayOfStringAttributedValuesType) SetForMarshal() {
	I.XMLName.Local = "t:Initials"
}

func (I *InitialsArrayOfStringAttributedValuesType) GetSchema() *Schema {
	return &SchemaTypes
}
