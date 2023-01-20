package elements

// The AttachmentId element identifies an item or file attachment. This element is used in CreateAttachment responses.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/attachmentid
import "encoding/xml"

type AttachmentId struct {
	XMLName xml.Name

	// Identifies the unique identifier of the attachment.
	Id *string `xml:"Id,attr"`
	// Identifies the change key of the root store item to which the attachment is attached.
	RootItemChangeKey *string `xml:"RootItemChangeKey,attr"`
	// Identifies the unique identifier of the root store item to which the attachment is attached.
	RootItemId *string `xml:"RootItemId,attr"`
}

func (A *AttachmentId) SetForMarshal() {
	A.XMLName.Local = "t:AttachmentId"
}

func (A *AttachmentId) GetSchema() *Schema {
	return &SchemaTypes
}
