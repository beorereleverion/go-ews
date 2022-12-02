package elements

// The Permissions element contains the collection of permissions for a folder.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/permissions
type Permissions struct {
	// The Permission element defines the access that a user has to a folder.
	Permission *Permission `xml:"t:Permission"`
}
