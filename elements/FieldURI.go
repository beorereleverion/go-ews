package elements

// The FieldURI element identifies frequently referenced properties by URI.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/fielduri
import "encoding/xml"

type FieldURI struct {
	XMLName xml.Name

	// Identifies the URI of the property.
	FieldURI *string `xml:"FieldURI,attr"`
}

func (F *FieldURI) SetForMarshal() {
	F.XMLName.Local = "t:FieldURI"
}

func (F *FieldURI) GetSchema() *Schema {
	return &SchemaTypes
}
