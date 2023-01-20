package elements

// The YomiFirstNames element specifies an array of phonetic Japanese first names and the identifiers of their source attributions for the associated persona.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/yomifirstnames
import "encoding/xml"

type YomiFirstNames struct {
	XMLName xml.Name

	// The StringAttributedValue element specifies an instance in an array of attributes associated with a persona element.
	StringAttributedValue *StringAttributedValue `xml:"StringAttributedValue"`
}

func (Y *YomiFirstNames) SetForMarshal() {
	Y.XMLName.Local = "t:YomiFirstNames"
}

func (Y *YomiFirstNames) GetSchema() *Schema {
	return &SchemaTypes
}
