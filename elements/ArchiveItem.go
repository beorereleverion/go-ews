package elements

// The ArchiveItem element contains the source folder Id and an array of item Ids for the associated archive item.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/archiveitem
type ArchiveItem struct {
	// The ArchiveSourceFolderId element specifies the Id of the source folder for the archive item.
	ArchiveSourceFolderId *ArchiveSourceFolderId `xml:"m:ArchiveSourceFolderId"`
	// The ItemIds element contains the unique identities of items, occurrence items, and recurring master items that are used to delete, send, get, move, or copy items in the Exchange store.
	ItemIds *ItemIds `xml:"m:ItemIds"`
}
