package elements

// The CreateAttachmentResponse element defines a response to a CreateAttachment request.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/createattachmentresponse
import "encoding/xml"

type CreateAttachmentResponse struct {
	XMLName xml.Name

	// The ResponseMessages element contains the response messages for an Exchange Web Services request.
	ResponseMessages *ResponseMessages `xml:"ResponseMessages"`
}

func (C *CreateAttachmentResponse) SetForMarshal() {
	C.XMLName.Local = "m:CreateAttachmentResponse"
}

func (C *CreateAttachmentResponse) GetSchema() *Schema {
	return &SchemaMessages
}
