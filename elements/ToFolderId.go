package elements

// The ToFolderId element represents the destination folder for a copied or moved item or folder.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/tofolderid
type ToFolderId struct {
	// The DistinguishedFolderId element identifies folders that can be referenced by name. If you do not use this element, you must use the FolderId element to identify a folder.
	DistinguishedFolderId *DistinguishedFolderId `xml:"t:DistinguishedFolderId"`
	// The FolderId element contains the identifier and change key of a folder.
	FolderId *FolderId `xml:"t:FolderId"`
}
