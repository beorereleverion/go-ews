package elements

// The SyncFolderItems element defines a request to synchronize items in an Exchange store folder.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/syncfolderitems
type SyncFolderItems struct {
	// The Ignore element identifies items to skip during synchronization.
	Ignore *Ignore `xml:"m:Ignore"`
	// The ItemShape element identifies a set of properties to return in a GetItem operation, FindItem operation, or SyncFolderItems operation response.
	ItemShape *ItemShape `xml:"m:ItemShape"`
	// The MaxChangesReturned element describes the maximum number of changes that can be returned in a synchronization response.
	MaxChangesReturned *MaxChangesReturned `xml:"t:MaxChangesReturned"`
	// The SyncFolderId element represents the folder that contains the items to synchronize.
	SyncFolderId *SyncFolderId `xml:"m:SyncFolderId"`
	// The SyncScope element specifies whether just items or items and folder associated information are returned in a synchronization response.
	SyncScope *SyncScope `xml:"m:SyncScope"`
	// The SyncState element contains a base64-encoded form of the synchronization data that is updated after each successful request. This is used to identify the synchronization state.
	SyncState *SyncState `xml:"m:SyncState"`
}
