package elements

// The ResponseMessages element represents a list of mail tips response messages.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/responsemessages-arrayofmailtipsresponsemessagetype
import "encoding/xml"

type ResponseMessagesArrayOfMailTipsResponseMessageType struct {
	XMLName xml.Name

	// The MailTipsResponseMessageType element represents mail tips settings.
	MailTipsResponseMessageType *MailTipsResponseMessageType `xml:"MailTipsResponseMessageType"`
}

func (R *ResponseMessagesArrayOfMailTipsResponseMessageType) SetForMarshal() {
	R.XMLName.Local = "m:ResponseMessages"
}

func (R *ResponseMessagesArrayOfMailTipsResponseMessageType) GetSchema() *Schema {
	return &SchemaMessages
}
