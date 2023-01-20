package elements

// The MailboxHoldResult element contains the result of the GetHoldOnMailboxes request.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/mailboxholdresult
import "encoding/xml"

type MailboxHoldResult struct {
	XMLName xml.Name

	// The HoldId element contains the mailbox hold identifier.
	HoldId *HoldId `xml:"HoldId"`
	// The MailboxHoldStatuses element specifies a list of one or more MailboxHoldStatus elements.
	MailboxHoldStatuses *MailboxHoldStatuses `xml:"MailboxHoldStatuses"`
	// The Query element contains the search query for the hold.
	Query *Query `xml:"Query"`
}

func (M *MailboxHoldResult) SetForMarshal() {
	M.XMLName.Local = "m:MailboxHoldResult"
}

func (M *MailboxHoldResult) GetSchema() *Schema {
	return &SchemaMessages
}
