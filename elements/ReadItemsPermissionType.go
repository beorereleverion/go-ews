package elements

// The ReadItems element indicates whether a user has permission to read items within a folder. This element was introduced in Microsoft Exchange Server 2007 Service Pack 1 (SP1).
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/readitems-permissiontype
type ReadItemsPermissionType string

const (
	// FullDetails
	ReadItemsPermissionTypeFullDetails ReadItemsPermissionType = `FullDetails`
	// None
	ReadItemsPermissionTypeNone ReadItemsPermissionType = `None`
)
