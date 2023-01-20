package elements

// The SetHoldOnMailboxes element contains a SetHoldOnMailboxes request.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/setholdonmailboxes
import "encoding/xml"

type SetHoldOnMailboxes struct {
	XMLName xml.Name

	// The ActionType element indicates the type of action for the hold.
	ActionType *ActionTypeHoldActionType `xml:"ActionType"`
	// The Deduplication element indicates whether the search result should remove duplicate items.
	Deduplication *Deduplication `xml:"Deduplication"`
	// The HoldId element contains the mailbox hold identifier.
	HoldId *HoldId `xml:"HoldId"`
	// The InPlaceHoldIdentity element specifies the identity of a hold that preserves the mailbox items.
	InPlaceHoldIdentity *InPlaceHoldIdentity `xml:"InPlaceHoldIdentity"`
	// The IncludeNonIndexableItems element contains a Boolean value to indicate whether to include items that cannot be indexed.
	IncludeNonIndexableItems *IncludeNonIndexableItems `xml:"IncludeNonIndexableItems"`
	// The Language element contains the language used for the search query.
	Language *Language `xml:"Language"`
	// The Mailboxes element contains a list of mailboxes affected by the hold.
	Mailboxes *MailboxesArrayOfStringsType `xml:"Mailboxes"`
	// The Query element contains the search query for the hold.
	Query *Query `xml:"Query"`
}

func (S *SetHoldOnMailboxes) SetForMarshal() {
	S.XMLName.Local = "m:SetHoldOnMailboxes"
}

func (S *SetHoldOnMailboxes) GetSchema() *Schema {
	return &SchemaMessages
}
