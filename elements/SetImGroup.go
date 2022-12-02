package elements

// The SetImGroup element represents a request to change the display name of an instant messaging group.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/setimgroup
type SetImGroup struct {
	// The GroupId element uniquely identifies a group.
	GroupId *GroupId `xml:"m:GroupId"`
	// The NewDisplayName element contains the updated display name of an instant messaging group.
	NewDisplayName *NewDisplayName `xml:"m:NewDisplayName"`
}
