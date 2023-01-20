package elements

// The SpouseNames element specifies an array of spouse or partner names and the identifiers of their source attributions for the associated persona.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/spousenames
import "encoding/xml"

type SpouseNames struct {
	XMLName xml.Name

	// The StringAttributedValue element specifies an instance in an array of attributes associated with a persona element.
	StringAttributedValue *StringAttributedValue `xml:"StringAttributedValue"`
}

func (S *SpouseNames) SetForMarshal() {
	S.XMLName.Local = "t:SpouseNames"
}

func (S *SpouseNames) GetSchema() *Schema {
	return &SchemaTypes
}
