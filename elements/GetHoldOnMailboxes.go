package elements

// The GetHoldOnMailboxes element indicates the beginning of the GetHoldOnMailboxes request.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/getholdonmailboxes
type GetHoldOnMailboxes struct {
	// The HoldId element contains the mailbox hold identifier.
	HoldId *HoldId `xml:"t:HoldId"`
}
