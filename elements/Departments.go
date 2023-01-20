package elements

// The Departments element specifies an array of department names and the identifiers of their source attributions for the associated persona.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/departments
import "encoding/xml"

type Departments struct {
	XMLName xml.Name

	// The StringAttributedValue element specifies an instance in an array of attributes associated with a persona element.
	StringAttributedValue *StringAttributedValue `xml:"StringAttributedValue"`
}

func (D *Departments) SetForMarshal() {
	D.XMLName.Local = "t:Departments"
}

func (D *Departments) GetSchema() *Schema {
	return &SchemaTypes
}
