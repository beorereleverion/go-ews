package elements

// The GetAttachmentResponse element defines a response to a GetAttachment request.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/getattachmentresponse
import "encoding/xml"

type GetAttachmentResponse struct {
	XMLName xml.Name

	// The ResponseMessages element contains the response messages for an Exchange Web Services request.
	ResponseMessages *ResponseMessages `xml:"ResponseMessages"`
}

func (G *GetAttachmentResponse) SetForMarshal() {
	G.XMLName.Local = "m:GetAttachmentResponse"
}

func (G *GetAttachmentResponse) GetSchema() *Schema {
	return &SchemaMessages
}
