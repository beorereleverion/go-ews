package elements

// The ReplyToItem element contains a reply to the sender of an item in the Exchange store.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/replytoitem
import "encoding/xml"

type ReplyToItem struct {
	XMLName xml.Name

	// The BccRecipients element represents a collection of recipients to receive a blind carbon copy (Bcc) of an e-mail message.
	BccRecipients *BccRecipients `xml:"BccRecipients"`
	// The Body element specifies the body of an item.
	Body *Body `xml:"Body"`
	// The CcRecipients element represents a collection of recipients that will receive a copy of the message.
	CcRecipients *CcRecipients `xml:"CcRecipients"`
	// The From element represents the address from which the message was sent.
	From *From `xml:"From"`
	// The IsDeliveryReceiptRequested element indicates whether the sender of an item requests a delivery receipt.
	IsDeliveryReceiptRequested *IsDeliveryReceiptRequested `xml:"IsDeliveryReceiptRequested"`
	// The IsReadReceiptRequested element indicates whether the sender of an item requests a read receipt.
	IsReadReceiptRequested *IsReadReceiptRequested `xml:"IsReadReceiptRequested"`
	// The NewBodyContent element represents the new body content of a message.
	NewBodyContent *NewBodyContent `xml:"NewBodyContent"`
	// The ReceivedBy element identifies the delegate in a delegate access scenario.
	ReceivedBy *ReceivedBy `xml:"ReceivedBy"`
	// The ReceivedRepresenting element identifies the principal in a delegate access scenario.
	ReceivedRepresenting *ReceivedRepresenting `xml:"ReceivedRepresenting"`
	// The ReferenceItemId element identifies the item to which the response object refers.
	ReferenceItemId *ReferenceItemId `xml:"ReferenceItemId"`
	// The Subject element represents the subject property of Exchange store items. The subject is limited to 255 characters.
	Subject *Subject `xml:"Subject"`
	// The ToRecipients element contains an array of recipients of an item. These are the primary recipients of an item.
	ToRecipients []*ToRecipients `xml:"ToRecipients"`
}

func (R *ReplyToItem) SetForMarshal() {
	R.XMLName.Local = "t:ReplyToItem"
}

func (R *ReplyToItem) GetSchema() *Schema {
	return &SchemaTypes
}
