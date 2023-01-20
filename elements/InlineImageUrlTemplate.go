package elements

// The InlineImageUrlTemplate element specifies a template for an inline image URL.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/inlineimageurltemplate
import "encoding/xml"

type InlineImageUrlTemplate struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (I *InlineImageUrlTemplate) SetForMarshal() {
	I.XMLName.Local = "t:InlineImageUrlTemplate"
}

func (I *InlineImageUrlTemplate) GetSchema() *Schema {
	return &SchemaTypes
}
