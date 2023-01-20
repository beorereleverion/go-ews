package elements

// The UpdateInboxRules element defines a request to update the Inbox rules in a mailbox in the server store.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/updateinboxrules
import "encoding/xml"

type UpdateInboxRules struct {
	XMLName xml.Name

	// The MailboxSmtpAddress element represents the SMTP address of the user whose Inbox rules are to be retrieved or updated; or whose password expiration date is to be retrieved.
	MailboxSmtpAddress *MailboxSmtpAddress `xml:"MailboxSmtpAddress"`
	// The Operations element contains an array of rule operations that can be performed on an Inbox.
	Operations *Operations `xml:"Operations"`
	// The RemoveOutlookRuleBlob element indicates whether to remove the Microsoft Outlook rule blob.
	RemoveOutlookRuleBlob *RemoveOutlookRuleBlob `xml:"RemoveOutlookRuleBlob"`
}

func (U *UpdateInboxRules) SetForMarshal() {
	U.XMLName.Local = "m:UpdateInboxRules"
}

func (U *UpdateInboxRules) GetSchema() *Schema {
	return &SchemaMessages
}
