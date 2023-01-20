package elements

// The AttachmentId element identifies a single attachment.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/attachmentid-getattachment-and-deleteattachment
import "encoding/xml"

type AttachmentIdGetAttachmentandDeleteAttachment struct {
	XMLName xml.Name

	// Specifies the attachment identifier.
	Id *string `xml:"Id,attr"`
}

func (A *AttachmentIdGetAttachmentandDeleteAttachment) SetForMarshal() {
	A.XMLName.Local = "t:AttachmentId"
}

func (A *AttachmentIdGetAttachmentandDeleteAttachment) GetSchema() *Schema {
	return &SchemaTypes
}
