package elements

// The MemberCorrelationKey element specifies the identifiers of the contacts that are part of the instant messaging (IM) group.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/membercorrelationkey
type MemberCorrelationKey struct {
	// The ItemId element contains the unique identifier and change key of an item in the Exchange store.
	ItemId *ItemId `xml:"t:ItemId"`
}
