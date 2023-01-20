package elements

// The DeleteAttachmentResponse defines a response to a DeleteAttachment request.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/deleteattachmentresponse
import "encoding/xml"

type DeleteAttachmentResponse struct {
	XMLName xml.Name

	// The ResponseMessages element contains the response messages for an Exchange Web Services request.
	ResponseMessages *ResponseMessages `xml:"ResponseMessages"`
}

func (D *DeleteAttachmentResponse) SetForMarshal() {
	D.XMLName.Local = "m:DeleteAttachmentResponse"
}

func (D *DeleteAttachmentResponse) GetSchema() *Schema {
	return &SchemaMessages
}
