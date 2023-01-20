package elements

// The GetClientIntentResponseMessage element is intended for internal use only.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/getclientintentresponsemessage
import "encoding/xml"

type GetClientIntentResponseMessage struct {
	XMLName xml.Name
	TEXT    interface{} `xml:",chardata"`
}

func (G *GetClientIntentResponseMessage) SetForMarshal() {
	G.XMLName.Local = "m:GetClientIntentResponseMessage"
}

func (G *GetClientIntentResponseMessage) GetSchema() *Schema {
	return &SchemaMessages
}
