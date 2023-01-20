package elements

// The GetClientIntent element is intended for internal use only.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/getclientintent
import "encoding/xml"

type GetClientIntent struct {
	XMLName xml.Name
	TEXT    interface{} `xml:",chardata"`
}

func (G *GetClientIntent) SetForMarshal() {
	G.XMLName.Local = "m:GetClientIntent"
}

func (G *GetClientIntent) GetSchema() *Schema {
	return &SchemaMessages
}
