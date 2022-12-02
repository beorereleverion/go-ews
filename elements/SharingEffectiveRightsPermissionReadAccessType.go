package elements

// The SharingEffectiveRights element indicates the permissions that the user has for the contact data that is being shared.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/sharingeffectiverights-permissionreadaccesstype
type SharingEffectiveRightsPermissionReadAccessType string

const (
	// FullDetails
	SharingEffectiveRightsPermissionReadAccessTypeFullDetails SharingEffectiveRightsPermissionReadAccessType = `FullDetails`
	// None
	SharingEffectiveRightsPermissionReadAccessTypeNone SharingEffectiveRightsPermissionReadAccessType = `None`
)
