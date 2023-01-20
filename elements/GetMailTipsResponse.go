package elements

// The GetMailTipsResponse element represents the response message for a GetMailTips operation.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/getmailtipsresponse
import "encoding/xml"

type GetMailTipsResponse struct {
	XMLName xml.Name

	// The ResponseMessages element represents a list of mail tips response messages.
	ResponseMessages *ResponseMessagesArrayOfMailTipsResponseMessageType `xml:"ResponseMessages"`
}

func (G *GetMailTipsResponse) SetForMarshal() {
	G.XMLName.Local = "m:GetMailTipsResponse"
}

func (G *GetMailTipsResponse) GetSchema() *Schema {
	return &SchemaMessages
}
