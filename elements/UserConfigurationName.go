package elements

// The UserConfigurationName element represents the name of a user configuration object. The user configuration object name is the identifier for a user configuration object.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/userconfigurationname
type UserConfigurationName struct {
	// The DistinguishedFolderId element identifies folders that can be referenced by name. If you do not use this element, you must use the FolderId element to identify a folder.
	DistinguishedFolderId *DistinguishedFolderId `xml:"t:DistinguishedFolderId"`
	// The FolderId element contains the identifier and change key of a folder.
	FolderId *FolderId `xml:"t:FolderId"`
	// Contains a string value that represents the name of a user configuration object. This attribute is required.
	Name *string `xml:"Name,attr"`
}
