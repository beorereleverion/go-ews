package elements

// The RestrictedGroupSids element represents a collection of restricted groups from a user's token.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/restrictedgroupsids
type RestrictedGroupSids struct {
	// The RestrictGroupIdentifier element represents the group security identifier (SID) and attributes for a restricted group within a user token.
	RestrictedGroupIdentifier *RestrictedGroupIdentifier `xml:"t:RestrictedGroupIdentifier"`
}
