package elements

// The FederatedDeliveryMailbox element represents the mailbox to which a cross-premise message was sent.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/federateddeliverymailbox
type FederatedDeliveryMailbox struct {
	// The EmailAddress element defines the primary SMTP address of a mailbox user.
	EmailAddress *EmailAddressNonEmptyStringType `xml:"t:EmailAddress"`
	// The ItemId element contains the unique identifier and change key of an item in the Exchange store.
	ItemId *ItemId `xml:"t:ItemId"`
	// The MailboxType element represents the type of mailbox that is represented by the e-mail address.
	MailboxType *MailboxType `xml:"t:MailboxType"`
	// The Name element represents the name of a mailbox user.
	Name *NameEmailAddressType `xml:"t:Name"`
	// The RoutingType element represents the routing protocol for the recipient.
	RoutingType *RoutingTypeEmailAddress `xml:"t:RoutingType"`
}
