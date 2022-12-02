package elements

// The RemoveDistributionGroupFromImList element represents a request to remove a specific instant messaging distribution list group.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/removedistributiongroupfromimlist
type RemoveDistributionGroupFromImList struct {
	// The GroupId element uniquely identifies a group.
	GroupId *GroupId `xml:"m:GroupId"`
}
