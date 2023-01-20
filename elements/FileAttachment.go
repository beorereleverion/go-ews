package elements

// The FileAttachment element represents a file that is attached to an item in the Exchange store.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/fileattachment
import "encoding/xml"

type FileAttachment struct {
	XMLName xml.Name

	// The AttachmentId element identifies an item or file attachment. This element is used in CreateAttachment responses.
	AttachmentId *AttachmentId `xml:"AttachmentId"`
	// The Content element contains the Base64-encoded contents of a file attachment.
	Content *Content `xml:"Content"`
	// The ContentId element represents an identifier for the contents of an attachment. ContentId can be set to any string value. Applications can use ContentId to implement their own identification mechanisms.
	ContentId *ContentId `xml:"ContentId"`
	// The ContentLocation element contains the Uniform Resource Identifier (URI) that corresponds to the location of the content of an attachment.
	ContentLocation *ContentLocation `xml:"ContentLocation"`
	// The ContentType element describes the Multipurpose Internet Mail Extensions (MIME) type of the attachment content.
	ContentType *ContentType `xml:"ContentType"`
	// The IsContactPhoto element indicates whether the file attachment is a contact picture.
	IsContactPhoto *IsContactPhoto `xml:"IsContactPhoto"`
	// The IsInline element represents whether the attachment appears inline within an item.
	IsInline *IsInline `xml:"IsInline"`
	// The LastModifiedTime element indicates when an item was last modified. This element is read-only.
	LastModifiedTime *LastModifiedTime `xml:"LastModifiedTime"`
	// The Name element represents the name of the attachment.
	Name *NameAttachmentType `xml:"Name"`
	// The Size element represents the size in bytes of an item or all the items in a conversation in the current folder. This property is read-only.
	Size *Size `xml:"Size"`
}

func (F *FileAttachment) SetForMarshal() {
	F.XMLName.Local = "t:FileAttachment"
}

func (F *FileAttachment) GetSchema() *Schema {
	return &SchemaTypes
}
