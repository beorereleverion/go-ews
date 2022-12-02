package elements

// The RemoveImGroup element represents a request to remove an instant messaging group.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/removeimgroup
type RemoveImGroup struct {
	// The GroupId element uniquely identifies a group.
	GroupId *GroupId `xml:"m:GroupId"`
}
