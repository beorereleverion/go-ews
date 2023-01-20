package elements

// The AttachmentIds element contains an array of attachment identifiers.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/attachmentids
import "encoding/xml"

type AttachmentIds struct {
	XMLName xml.Name

	// The AttachmentId element identifies a single attachment.
	AttachmentId *AttachmentIdGetAttachmentandDeleteAttachment `xml:"AttachmentId"`
}

func (A *AttachmentIds) SetForMarshal() {
	A.XMLName.Local = "m:AttachmentIds"
}

func (A *AttachmentIds) GetSchema() *Schema {
	return &SchemaMessages
}
