package elements

// The SubscribeResponse element defines a response to a Subscribe request.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/subscriberesponse
import "encoding/xml"

type SubscribeResponse struct {
	XMLName xml.Name

	// The ResponseMessages element contains the response messages for an Exchange Web Services request.
	ResponseMessages *ResponseMessages `xml:"ResponseMessages"`
}

func (S *SubscribeResponse) SetForMarshal() {
	S.XMLName.Local = "m:SubscribeResponse"
}

func (S *SubscribeResponse) GetSchema() *Schema {
	return &SchemaMessages
}
