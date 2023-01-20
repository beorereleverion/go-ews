package elements

// The BodyContentAttributedValue element specifies the body content of an item.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/bodycontentattributedvalue
import "encoding/xml"

type BodyContentAttributedValue struct {
	XMLName xml.Name

	// The Attributions element specifies an array of attribution information for one or more of the contacts or Active Directory recipients aggregated into the associated persona.
	Attributions *AttributionsArrayOfPersonaAttributionsType `xml:"Attributions"`
	// The Value element specifies the value of a BodyContentAttributedValue element.
	Value *ValueBodyContentType `xml:"Value"`
}

func (B *BodyContentAttributedValue) SetForMarshal() {
	B.XMLName.Local = "t:BodyContentAttributedValue"
}

func (B *BodyContentAttributedValue) GetSchema() *Schema {
	return &SchemaTypes
}
