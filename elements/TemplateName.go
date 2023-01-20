package elements

// The TemplateName element specifies the name of the Rights Management template.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/templatename
import "encoding/xml"

type TemplateName struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (T *TemplateName) SetForMarshal() {
	T.XMLName.Local = "t:TemplateName"
}

func (T *TemplateName) GetSchema() *Schema {
	return &SchemaTypes
}
