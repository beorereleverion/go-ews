package elements

// The GetHoldOnMailboxes element contains the request to get the hold status for a mailbox.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/getholdonmailboxesresponse
type GetHoldOnMailboxesResponse struct {
	// The HoldId element contains the mailbox hold identifier.
	HoldId *HoldId `xml:"t:HoldId"`
}
