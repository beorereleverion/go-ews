package elements

// The GetUserOofSettingsRequest element is the root element that contains the arguments used to get a mailbox user's Out of Office (OOF) settings.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/getuseroofsettingsrequest
type GetUserOofSettingsRequest struct {
	// The Mailbox element represents the mailbox user for a SetUserOofSettings or GetUserOofSettings request.
	Mailbox *MailboxAvailability `xml:"t:Mailbox"`
}
