package elements

// The Mailbox element identifies a mail-enabled Active Directory object.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/mailbox
type Mailbox struct {
	// Defines the name of the mailbox user. This element is optional.
	Name string `xml:"Name"`
	// Defines the Simple Mail Transfer Protocol (SMTP) address of a mailbox user. This element is optional.
	EmailAddress string `xml:"EmailAddress"`
	// Defines the routing that is used for the mailbox. The default is SMTP. This element is optional.
	RoutingType string `xml:"RoutingType"`
	// Defines the mailbox type of a mailbox user. This element is optional.
	MailboxType MailboxType `xml:"t:MailboxType"`
	// Defines the item identifier of a contact or private distribution list for recipients from a user's contacts folder. This element is optional.
	ItemId ItemId `xml:"t:ItemId"`
}
