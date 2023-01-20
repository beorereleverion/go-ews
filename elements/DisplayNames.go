package elements

// The DisplayNames element specifies an array of display names and the identifiers of their source attributions for the associated persona.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/displaynames
import "encoding/xml"

type DisplayNames struct {
	XMLName xml.Name

	// The StringAttributedValue element specifies an instance in an array of attributes associated with a persona element.
	StringAttributedValue *StringAttributedValue `xml:"StringAttributedValue"`
}

func (D *DisplayNames) SetForMarshal() {
	D.XMLName.Local = "t:DisplayNames"
}

func (D *DisplayNames) GetSchema() *Schema {
	return &SchemaTypes
}
