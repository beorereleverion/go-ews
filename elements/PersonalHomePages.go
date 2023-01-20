package elements

// The PersonalHomePages element specifies an array of home pages and the identifiers of their source attributions for the associated persona.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/personalhomepages
import "encoding/xml"

type PersonalHomePages struct {
	XMLName xml.Name

	// The StringAttributedValue element specifies an instance in an array of attributes associated with a persona element.
	StringAttributedValue *StringAttributedValue `xml:"StringAttributedValue"`
}

func (P *PersonalHomePages) SetForMarshal() {
	P.XMLName.Local = "t:PersonalHomePages"
}

func (P *PersonalHomePages) GetSchema() *Schema {
	return &SchemaTypes
}
