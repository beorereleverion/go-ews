package elements

// The CalendarFolder element represents a folder that primarily contains calendar items.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/calendarfolder
type CalendarFolder struct {
	// The ChildFolderCount element represents the number of immediate child folders that are contained within a folder. This property is read-only.
	ChildFolderCount *ChildFolderCount `xml:"t:ChildFolderCount"`
	// The DisplayName element defines the display name of a folder, contact, distribution list, delegate user, location, or rule.
	DisplayName *DisplayNamestring `xml:"t:DisplayName"`
	// The EffectiveRights element contains the client's rights based on the permission settings for the item or folder. This element is read-only.
	EffectiveRights *EffectiveRights `xml:"t:EffectiveRights"`
	// The ExtendedProperty element identifies extended MAPI properties on folders and items.
	ExtendedProperty *ExtendedProperty `xml:"t:ExtendedProperty"`
	// The FolderClass element represents the folder class for a folder.
	FolderClass *FolderClass `xml:"t:FolderClass"`
	// The FolderId element contains the identifier and change key of a folder.
	FolderId *FolderId `xml:"t:FolderId"`
	// The ManagedFolderInformation element contains information about a managed custom folder.
	ManagedFolderInformation *ManagedFolderInformation `xml:"t:ManagedFolderInformation"`
	// The ParentFolderId element represents the identifier of the parent folder that contains the item or folder.
	ParentFolderId *ParentFolderId `xml:"t:ParentFolderId"`
	// The PermissionSet element contains all the permissions that are configured for a calendar folder.
	PermissionSet *PermissionSetCalendarPermissionSetType `xml:"t:PermissionSet"`
	// The SharingEffectiveRights element indicates the permissions that the user has for the calendar data that is being shared.
	SharingEffectiveRights *SharingEffectiveRightsCalendarPermissionReadAccessType `xml:"t:SharingEffectiveRights"`
	// The TotalCount element represents the total count of items within a given folder.
	TotalCount *TotalCount `xml:"t:TotalCount"`
}
