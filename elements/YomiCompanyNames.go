package elements

// The YomiCompanyNames element specifies an array of phonetic Japanese company names and the identifiers of their source attributions for the associated persona.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/yomicompanynames
import "encoding/xml"

type YomiCompanyNames struct {
	XMLName xml.Name

	// The StringAttributedValue element specifies an instance in an array of attributes associated with a persona element.
	StringAttributedValue *StringAttributedValue `xml:"StringAttributedValue"`
}

func (Y *YomiCompanyNames) SetForMarshal() {
	Y.XMLName.Local = "t:YomiCompanyNames"
}

func (Y *YomiCompanyNames) GetSchema() *Schema {
	return &SchemaTypes
}
