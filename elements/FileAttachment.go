package elements

// The FileAttachment element represents a file that is attached to an item in the Exchange store.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/fileattachment
type FileAttachment struct {
	// The AttachmentId element identifies an item or file attachment. This element is used in CreateAttachment responses.
	AttachmentId *AttachmentId `xml:"t:AttachmentId"`
	// The Content element contains the Base64-encoded contents of a file attachment.
	Content *Content `xml:"t:Content"`
	// The ContentId element represents an identifier for the contents of an attachment. ContentId can be set to any string value. Applications can use ContentId to implement their own identification mechanisms.
	ContentId *ContentId `xml:"t:ContentId"`
	// The ContentLocation element contains the Uniform Resource Identifier (URI) that corresponds to the location of the content of an attachment.
	ContentLocation *ContentLocation `xml:"t:ContentLocation"`
	// The ContentType element describes the Multipurpose Internet Mail Extensions (MIME) type of the attachment content.
	ContentType *ContentType `xml:"t:ContentType"`
	// The IsContactPhoto element indicates whether the file attachment is a contact picture.
	IsContactPhoto *IsContactPhoto `xml:"t:IsContactPhoto"`
	// The IsInline element represents whether the attachment appears inline within an item.
	IsInline *IsInline `xml:"t:IsInline"`
	// The LastModifiedTime element indicates when an item was last modified. This element is read-only.
	LastModifiedTime *LastModifiedTime `xml:"t:LastModifiedTime"`
	// The Name element represents the name of the attachment.
	Name *NameAttachmentType `xml:"t:Name"`
	// The Size element represents the size in bytes of an item or all the items in a conversation in the current folder. This property is read-only.
	Size *Size `xml:"t:Size"`
}
