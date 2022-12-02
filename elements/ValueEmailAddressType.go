package elements

// The Value element specifies the value of an EmailAddress associated with an attributions array.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/value-emailaddresstype
type ValueEmailAddressType struct {
	// The EmailAddress element defines the primary SMTP address of a mailbox user.
	EmailAddress *EmailAddressNonEmptyStringType `xml:"t:EmailAddress"`
	// The ItemId element contains the unique identifier and change key of an item in the Exchange store.
	ItemId *ItemId `xml:"t:ItemId"`
	// The MailboxType element represents the type of mailbox that is represented by the e-mail address.
	MailboxType *MailboxType `xml:"t:MailboxType"`
	// The Name element specifies a search refiner name.
	Name *Namestring `xml:"t:Name"`
	// The OriginalDisplayName element contains the original display name associated with an email address.
	OriginalDisplayName *OriginalDisplayName `xml:"t:OriginalDisplayName"`
	// The RoutingType element defines the address type for the mailbox.
	RoutingType *RoutingTypeEmailAddressType `xml:"t:RoutingType"`
}
