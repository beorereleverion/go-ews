package elements

// The ServerReplyWithMessage element indicates the ID of the template message that is to be sent as a reply to incoming messages.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/serverreplywithmessage
import "encoding/xml"

type ServerReplyWithMessage struct {
	XMLName xml.Name

	// The ItemId element contains the unique identifier and change key of an item in the Exchange store.
	ItemId *ItemId `xml:"ItemId"`
}

func (S *ServerReplyWithMessage) SetForMarshal() {
	S.XMLName.Local = "m:ServerReplyWithMessage"
}

func (S *ServerReplyWithMessage) GetSchema() *Schema {
	return &SchemaMessages
}
