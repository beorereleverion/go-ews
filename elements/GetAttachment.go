package elements

// The GetAttachment element is the root element in a request to get an attachment from the Exchange store.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/getattachment
import "encoding/xml"

type GetAttachment struct {
	XMLName xml.Name

	// The AttachmentIds element contains an array of attachment identifiers.
	AttachmentIds *AttachmentIds `xml:"AttachmentIds"`
	// The AttachmentShape element identifies additional properties to return in a response to a GetAttachment request.
	AttachmentShape *AttachmentShape `xml:"AttachmentShape"`
	// The GetAttachment element is the root element in a request to get an attachment from the Exchange store.
	GetAttachment *GetAttachment `xml:"GetAttachment"`
}

func (G *GetAttachment) SetForMarshal() {
	G.XMLName.Local = "m:GetAttachment"
}

func (G *GetAttachment) GetSchema() *Schema {
	return &SchemaMessages
}
