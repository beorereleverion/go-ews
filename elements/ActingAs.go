package elements

// The ActingAs element identifies who the caller is sending as.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/actingas
type ActingAs struct {
	// The EmailAddress element defines the primary SMTP address of a mailbox user.
	EmailAddress *EmailAddressNonEmptyStringType `xml:"t:EmailAddress"`
	// The RoutingType element represents the routing protocol for the recipient.
	RoutingType *RoutingTypeEmailAddress `xml:"t:RoutingType"`
}
