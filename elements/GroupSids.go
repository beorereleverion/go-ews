package elements

// The GroupSids element represents a collection of Active Directory directory service group object security identifiers.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/groupsids
type GroupSids struct {
	// The GroupIdentifier element represents a single security identifier and attribute for an Active Directory directory service object group of which the account is a member.
	GroupIdentifier *GroupIdentifier `xml:"t:GroupIdentifier"`
}
