package elements

// The SearchPreviewItem element specifies the item preview for a discovery search.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/searchpreviewitem
import "encoding/xml"

type SearchPreviewItem struct {
	XMLName xml.Name

	// The BccRecipients element represents a collection of recipients to receive a blind carbon copy (Bcc) of an e-mail message.
	BccRecipients *BccRecipients `xml:"BccRecipients"`
	// The CcRecipients element represents a collection of recipients that will receive a copy of the message.
	CcRecipients *CcRecipients `xml:"CcRecipients"`
	// The CreatedTime element specifies the time at which the item was created.
	CreatedTime *CreatedTime `xml:"CreatedTime"`
	// The ExtendedProperties element specifies an array of additional properties.
	ExtendedProperties *ExtendedPropertiesNonEmptyArrayOfExtendedPropertyType `xml:"ExtendedProperties"`
	// The HasAttachment element specifies a Boolean value to indicate whether the item has attachments.
	HasAttachment *HasAttachment `xml:"HasAttachment"`
	// The Id element specifies the identifier of an item.
	ID *IDItemIdType `xml:"ID"`
	// The Importance element describes the importance of an item or the aggregated importance of all items in a conversation in the current folder.
	Importance *Importance `xml:"Importance"`
	// The ItemClass element represents the message class of an item.
	ItemClass *ItemClass `xml:"ItemClass"`
	// The Mailbox element contains the mailbox identifier and the user's primary Simple Mail Transfer Protocol (SMTP) address.
	Mailbox *MailboxPreviewItemMailboxType `xml:"Mailbox"`
	// The OwaLink element specifies the link to preview an item in Microsoft Outlook Web App.
	OwaLink *OwaLink `xml:"OwaLink"`
	// The ParentId element specifies the identifier of the parent item in a search preview.
	ParentId *ParentId `xml:"ParentId"`
	// The Preview element specifies the first 256 characters of an item for display.
	Preview *Preview `xml:"Preview"`
	// The Read element indicates whether a client can read a folder or item. This element was introduced in Microsoft Exchange Server 2007 Service Pack 1 (SP1).
	Read *Read `xml:"Read"`
	// The ReceivedTime element specifies the time at which an item was received.
	ReceivedTime *ReceivedTime `xml:"ReceivedTime"`
	// The Sender element specifies the e-mail address of the person who sent an item.
	Sender *Senderstring `xml:"Sender"`
	// The SentTime element specifies the time at which the item was sent.
	SentTime *SentTime `xml:"SentTime"`
	// The Size element specifies the total size of one or more mailbox items.
	Size *Sizelong `xml:"Size"`
	// The SortValue element specifies a value used for sorting.
	SortValue *SortValue `xml:"SortValue"`
	// The Subject element represents the subject property of Exchange store items. The subject is limited to 255 characters.
	Subject *Subject `xml:"Subject"`
	// The ToRecipients element specifies a list of recipients to whom the item was sent.
	ToRecipients []*ToRecipientsArrayOfSmtpAddressType `xml:"ToRecipients"`
	// The UniqueHash element specifies a unique hash value used for de-duplication.
	UniqueHash *UniqueHash `xml:"UniqueHash"`
}

func (S *SearchPreviewItem) SetForMarshal() {
	S.XMLName.Local = "t:SearchPreviewItem"
}

func (S *SearchPreviewItem) GetSchema() *Schema {
	return &SchemaTypes
}
