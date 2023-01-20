package elements

// The GetInboxRules element defines a request to get the Inbox rules on a mailbox in the server store.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/getinboxrules
import "encoding/xml"

type GetInboxRules struct {
	XMLName xml.Name

	// The MailboxSmtpAddress element represents the SMTP address of the user whose Inbox rules are to be retrieved or updated; or whose password expiration date is to be retrieved.
	MailboxSmtpAddress *MailboxSmtpAddress `xml:"MailboxSmtpAddress"`
}

func (G *GetInboxRules) SetForMarshal() {
	G.XMLName.Local = "m:GetInboxRules"
}

func (G *GetInboxRules) GetSchema() *Schema {
	return &SchemaMessages
}
