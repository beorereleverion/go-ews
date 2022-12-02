package elements

// The ReplyToItem element contains a reply to the sender of an item in the Exchange store.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/replytoitem
type ReplyToItem struct {
	// The BccRecipients element represents a collection of recipients to receive a blind carbon copy (Bcc) of an e-mail message.
	BccRecipients *BccRecipients `xml:"t:BccRecipients"`
	// The Body element specifies the body of an item.
	Body *Body `xml:"t:Body"`
	// The CcRecipients element represents a collection of recipients that will receive a copy of the message.
	CcRecipients *CcRecipients `xml:"t:CcRecipients"`
	// The From element represents the address from which the message was sent.
	From *From `xml:"t:From"`
	// The IsDeliveryReceiptRequested element indicates whether the sender of an item requests a delivery receipt.
	IsDeliveryReceiptRequested *IsDeliveryReceiptRequested `xml:"t:IsDeliveryReceiptRequested"`
	// The IsReadReceiptRequested element indicates whether the sender of an item requests a read receipt.
	IsReadReceiptRequested *IsReadReceiptRequested `xml:"t:IsReadReceiptRequested"`
	// The NewBodyContent element represents the new body content of a message.
	NewBodyContent *NewBodyContent `xml:"t:NewBodyContent"`
	// The ReceivedBy element identifies the delegate in a delegate access scenario.
	ReceivedBy *ReceivedBy `xml:"t:ReceivedBy"`
	// The ReceivedRepresenting element identifies the principal in a delegate access scenario.
	ReceivedRepresenting *ReceivedRepresenting `xml:"t:ReceivedRepresenting"`
	// The ReferenceItemId element identifies the item to which the response object refers.
	ReferenceItemId *ReferenceItemId `xml:"t:ReferenceItemId"`
	// The Subject element represents the subject property of Exchange store items. The subject is limited to 255 characters.
	Subject *Subject `xml:"t:Subject"`
	// The ToRecipients element contains an array of recipients of an item. These are the primary recipients of an item.
	ToRecipients *ToRecipients `xml:"t:ToRecipients"`
}
