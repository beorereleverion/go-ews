package elements

// The SendNotification element contains the push notifications that are sent by the computer that is running Microsoft Exchange Server 2007 to the client application.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/sendnotification
import "encoding/xml"

type SendNotification struct {
	XMLName xml.Name

	// The ResponseMessages element contains the response messages for an Exchange Web Services request.
	ResponseMessages *ResponseMessages `xml:"ResponseMessages"`
}

func (S *SendNotification) SetForMarshal() {
	S.XMLName.Local = "m:SendNotification"
}

func (S *SendNotification) GetSchema() *Schema {
	return &SchemaMessages
}
