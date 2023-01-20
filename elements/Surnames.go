package elements

// The Surnames element specifies an array of surname values and the identifiers of their source attributions for the associated persona.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/surnames
import "encoding/xml"

type Surnames struct {
	XMLName xml.Name

	// The StringAttributedValue element specifies an instance in an array of attributes associated with a persona element.
	StringAttributedValue *StringAttributedValue `xml:"StringAttributedValue"`
}

func (S *Surnames) SetForMarshal() {
	S.XMLName.Local = "t:Surnames"
}

func (S *Surnames) GetSchema() *Schema {
	return &SchemaTypes
}
