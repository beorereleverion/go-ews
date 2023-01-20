package elements

// The CompanyNames element contains an array of company names and the identifiers of their source attributions for the associated persona.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/companynames
import "encoding/xml"

type CompanyNames struct {
	XMLName xml.Name

	// The StringAttributedValue element specifies an instance in an array of attributes associated with a persona element.
	StringAttributedValue *StringAttributedValue `xml:"StringAttributedValue"`
}

func (C *CompanyNames) SetForMarshal() {
	C.XMLName.Local = "t:CompanyNames"
}

func (C *CompanyNames) GetSchema() *Schema {
	return &SchemaTypes
}
