package elements

// The Email element represents the mailbox user for a GetUserAvailability query.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/email-emailaddresstype
type EmailEmailAddressType struct {
	// The Address element represents the e-mail address of the mailbox user.
	Address *Addressstring `xml:"t:Address"`
	// The Name element represents the display name of the mailbox user.
	Name *NameEmailAddress `xml:"t:Name"`
	// The RoutingType element represents the routing protocol for the recipient.
	RoutingType *RoutingTypeEmailAddress `xml:"t:RoutingType"`
}
