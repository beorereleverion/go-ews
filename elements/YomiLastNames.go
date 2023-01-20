package elements

// The YomiLastNames element specifies an array of phonetic Japanese last names and the identifiers of their source attributions for the associated persona.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/yomilastnames
import "encoding/xml"

type YomiLastNames struct {
	XMLName xml.Name

	// The StringAttributedValue element specifies an instance in an array of attributes associated with a persona element.
	StringAttributedValue *StringAttributedValue `xml:"StringAttributedValue"`
}

func (Y *YomiLastNames) SetForMarshal() {
	Y.XMLName.Local = "t:YomiLastNames"
}

func (Y *YomiLastNames) GetSchema() *Schema {
	return &SchemaTypes
}
