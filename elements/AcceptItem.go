package elements

// The AcceptItem element represents an Accept reply to a meeting request.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/acceptitem
import "encoding/xml"

type AcceptItem struct {
	XMLName xml.Name

	// The Attachments element contains the items or files that are attached to an item in the Exchange store.
	Attachments *Attachments `xml:"Attachments"`
	// The BccRecipients element represents a collection of recipients to receive a blind carbon copy (Bcc) of an e-mail message.
	BccRecipients *BccRecipients `xml:"BccRecipients"`
	// The Body element specifies the body of an item.
	Body *Body `xml:"Body"`
	// The CcRecipients element represents a collection of recipients that will receive a copy of the message.
	CcRecipients *CcRecipients `xml:"CcRecipients"`
	// The InternetMessageHeaders element contains a collection of some of the Internet message headers that are contained in an item in a mailbox. To get the entire collection of Internet message headers, use the PR_TRANSPORT_MESSAGE_HEADERS property. For more information about EWS and Internet message headers, seeGetting Internet message headersin EWS, MIME, and the missing Internet message headers.
	InternetMessageHeaders *InternetMessageHeaders `xml:"InternetMessageHeaders"`
	// The IsDeliveryReceiptRequested element indicates whether the sender of an item requests a delivery receipt.
	IsDeliveryReceiptRequested *IsDeliveryReceiptRequested `xml:"IsDeliveryReceiptRequested"`
	// The IsReadReceiptRequested element indicates whether the sender of an item requests a read receipt.
	IsReadReceiptRequested *IsReadReceiptRequested `xml:"IsReadReceiptRequested"`
	// The ItemClass element represents the message class of an item.
	ItemClass *ItemClass `xml:"ItemClass"`
	// The ProposedEnd element specifies the proposed end time of a meeting.
	ProposedEnd *ProposedEnd `xml:"ProposedEnd"`
	// The ProposedStart element specifies the proposed start time of a meeting.
	ProposedStart *ProposedStart `xml:"ProposedStart"`
	// The ReceivedBy element identifies the delegate in a delegate access scenario.
	ReceivedBy *ReceivedBy `xml:"ReceivedBy"`
	// The ReceivedRepresenting element identifies the principal in a delegate access scenario.
	ReceivedRepresenting *ReceivedRepresenting `xml:"ReceivedRepresenting"`
	// The ReferenceItemId element identifies the item to which the response object refers.
	ReferenceItemId *ReferenceItemId `xml:"ReferenceItemId"`
	// The Sender element identifies the sender of an item.
	Sender *Sender `xml:"Sender"`
	// The Sensitivity element indicates the sensitivity level of an item.
	Sensitivity *Sensitivity `xml:"Sensitivity"`
	// The ToRecipients element contains an array of recipients of an item. These are the primary recipients of an item.
	ToRecipients *ToRecipients `xml:"ToRecipients"`
}

func (A *AcceptItem) SetForMarshal() {
	A.XMLName.Local = "t:AcceptItem"
}

func (A *AcceptItem) GetSchema() *Schema {
	return &SchemaTypes
}
