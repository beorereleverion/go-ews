package elements

// The TentativelyAcceptItem element represents a Tentative reply to a meeting request.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/tentativelyacceptitem
type TentativelyAcceptItem struct {
	// The Attachments element contains the items or files that are attached to an item in the Exchange store.
	Attachments *Attachments `xml:"t:Attachments"`
	// The BccRecipients element represents a collection of recipients to receive a blind carbon copy (Bcc) of an e-mail message.
	BccRecipients *BccRecipients `xml:"t:BccRecipients"`
	// The Body element specifies the body of an item.
	Body *Body `xml:"t:Body"`
	// The CcRecipients element represents a collection of recipients that will receive a copy of the message.
	CcRecipients *CcRecipients `xml:"t:CcRecipients"`
	// The InternetMessageHeaders element contains a collection of some of the Internet message headers that are contained in an item in a mailbox. To get the entire collection of Internet message headers, use the PR_TRANSPORT_MESSAGE_HEADERS property. For more information about EWS and Internet message headers, seeGetting Internet message headersin EWS, MIME, and the missing Internet message headers.
	InternetMessageHeaders *InternetMessageHeaders `xml:"t:InternetMessageHeaders"`
	// The IsDeliveryReceiptRequested element indicates whether the sender of an item requests a delivery receipt.
	IsDeliveryReceiptRequested *IsDeliveryReceiptRequested `xml:"t:IsDeliveryReceiptRequested"`
	// The IsReadReceiptRequested element indicates whether the sender of an item requests a read receipt.
	IsReadReceiptRequested *IsReadReceiptRequested `xml:"t:IsReadReceiptRequested"`
	// The ItemClass element represents the message class of an item.
	ItemClass *ItemClass `xml:"t:ItemClass"`
	// The ProposedEnd element specifies the proposed end time of a meeting.
	ProposedEnd *ProposedEnd `xml:"t:ProposedEnd"`
	// The ProposedStart element specifies the proposed start time of a meeting.
	ProposedStart *ProposedStart `xml:"t:ProposedStart"`
	// The ReceivedBy element identifies the delegate in a delegate access scenario.
	ReceivedBy *ReceivedBy `xml:"t:ReceivedBy"`
	// The ReceivedRepresenting element identifies the principal in a delegate access scenario.
	ReceivedRepresenting *ReceivedRepresenting `xml:"t:ReceivedRepresenting"`
	// The ReferenceItemId element identifies the item to which the response object refers.
	ReferenceItemId *ReferenceItemId `xml:"t:ReferenceItemId"`
	// The Sender element identifies the sender of an item.
	Sender *Sender `xml:"t:Sender"`
	// The Sensitivity element indicates the sensitivity level of an item.
	Sensitivity *Sensitivity `xml:"t:Sensitivity"`
	// The ToRecipients element contains an array of recipients of an item. These are the primary recipients of an item.
	ToRecipients *ToRecipients `xml:"t:ToRecipients"`
}
