package elements

// The TemplateDescription element specifies the description of the Rights Management template.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/templatedescription
import "encoding/xml"

type TemplateDescription struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (T *TemplateDescription) SetForMarshal() {
	T.XMLName.Local = "t:TemplateDescription"
}

func (T *TemplateDescription) GetSchema() *Schema {
	return &SchemaTypes
}
