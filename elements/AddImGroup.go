package elements

// The AddImGroup element defines a request to add a new instant messaging group.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/addimgroup
type AddImGroup struct {
	// The DisplayName element defines the display name of a folder, contact, distribution list, delegate user, location, or rule.
	DisplayName *DisplayNamestring `xml:"t:DisplayName"`
}
