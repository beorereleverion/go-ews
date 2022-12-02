package elements

// The Members element provides the list of members for a distribution list.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/members-memberlisttype
type MembersMemberListType struct {
	// The Member element represents a member of a distribution list.
	Member *Member `xml:"t:Member"`
}
