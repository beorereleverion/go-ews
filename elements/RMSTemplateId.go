package elements

// The RMSTemplateId element specifies the identifier of the Rights Management template.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/rmstemplateid
import "encoding/xml"

type RMSTemplateId struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (R *RMSTemplateId) SetForMarshal() {
	R.XMLName.Local = "t:RMSTemplateId"
}

func (R *RMSTemplateId) GetSchema() *Schema {
	return &SchemaTypes
}
