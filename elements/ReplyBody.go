package elements

// The ReplyBody element contains an Out of Office (OOF) message and the language used for the message.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/replybody
import "encoding/xml"

type ReplyBody struct {
	XMLName xml.Name

	// The Message element contains the out of Office (OOF) response.
	Message *MessageAvailability `xml:"Message"`
}

func (R *ReplyBody) SetForMarshal() {
	R.XMLName.Local = "t:ReplyBody"
}

func (R *ReplyBody) GetSchema() *Schema {
	return &SchemaTypes
}
