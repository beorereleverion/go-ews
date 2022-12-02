package elements

// The FoldersToIgnore element identifies a list of folders that are ignored when getting items in a conversation. All conversation items in the ignored folders are not returned in a GetConversationItems response.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/folderstoignore
type FoldersToIgnore struct {
	// The DistinguishedFolderId element identifies folders that can be referenced by name. If you do not use this element, you must use the FolderId element to identify a folder.
	DistinguishedFolderId *DistinguishedFolderId `xml:"t:DistinguishedFolderId"`
	// The FolderId element contains the identifier and change key of a folder.
	FolderId *FolderId `xml:"t:FolderId"`
}
