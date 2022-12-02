package elements

// The ItemAttachment element represents an Exchange item that is attached to another Exchange item.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/itemattachment
type ItemAttachment struct {
	// The AttachmentId element identifies an item or file attachment. This element is used in CreateAttachment responses.
	AttachmentId *AttachmentId `xml:"t:AttachmentId"`
	// The CalendarItem element represents an Exchange calendar item.
	CalendarItem *CalendarItem `xml:"t:CalendarItem"`
	// The Contact element represents a contact item in the Exchange store.
	Contact *Contact `xml:"t:Contact"`
	// The ContentId element represents an identifier for the contents of an attachment. ContentId can be set to any string value. Applications can use ContentId to implement their own identification mechanisms.
	ContentId *ContentId `xml:"t:ContentId"`
	// The ContentLocation element contains the Uniform Resource Identifier (URI) that corresponds to the location of the content of an attachment.
	ContentLocation *ContentLocation `xml:"t:ContentLocation"`
	// The ContentType element describes the Multipurpose Internet Mail Extensions (MIME) type of the attachment content.
	ContentType *ContentType `xml:"t:ContentType"`
	// The IsInline element represents whether the attachment appears inline within an item.
	IsInline *IsInline `xml:"t:IsInline"`
	// The Item element represents a generic item in the Exchange store.
	Item *Item `xml:"t:Item"`
	// The LastModifiedTime element indicates when an item was last modified. This element is read-only.
	LastModifiedTime *LastModifiedTime `xml:"t:LastModifiedTime"`
	// The MeetingCancellation element represents a meeting cancellation in the Exchange store.
	MeetingCancellation *MeetingCancellation `xml:"t:MeetingCancellation"`
	// The MeetingMessage element represents a meeting in the Exchange store.
	MeetingMessage *MeetingMessage `xml:"t:MeetingMessage"`
	// The MeetingRequest element represents a meeting request in the Exchange store.
	MeetingRequest *MeetingRequest `xml:"t:MeetingRequest"`
	// The MeetingResponse element represents a meeting response in the Exchange store.
	MeetingResponse *MeetingResponse `xml:"t:MeetingResponse"`
	// The Message element represents a Microsoft Exchange e-mail message.
	Message *Message `xml:"t:Message"`
	// The Name element represents the name of the attachment.
	Name *NameAttachmentType `xml:"t:Name"`
	// The Size element represents the size in bytes of an item or all the items in a conversation in the current folder. This property is read-only.
	Size *Size `xml:"t:Size"`
	// The Task element represents a task in the Exchange store.
	Task *Task `xml:"t:Task"`
}
