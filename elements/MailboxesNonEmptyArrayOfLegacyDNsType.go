package elements

// The Mailboxes element specifies an array of mailboxes identified by legacy distinguished name.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/mailboxes-nonemptyarrayoflegacydnstype
type MailboxesNonEmptyArrayOfLegacyDNsType struct {
	// The LegacyDN element identifies a mailbox by its legacy distinguished name.
	LegacyDN *LegacyDN `xml:"t:LegacyDN"`
}
