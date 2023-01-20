package elements

// The SyncFolderItems element defines a request to synchronize items in an Exchange store folder.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/syncfolderitems
import "encoding/xml"

type SyncFolderItems struct {
	XMLName xml.Name

	// The Ignore element identifies items to skip during synchronization.
	Ignore *Ignore `xml:"Ignore"`
	// The ItemShape element identifies a set of properties to return in a GetItem operation, FindItem operation, or SyncFolderItems operation response.
	ItemShape *ItemShape `xml:"ItemShape"`
	// The MaxChangesReturned element describes the maximum number of changes that can be returned in a synchronization response.
	MaxChangesReturned *MaxChangesReturned `xml:"MaxChangesReturned"`
	// The SyncFolderId element represents the folder that contains the items to synchronize.
	SyncFolderId *SyncFolderId `xml:"SyncFolderId"`
	// The SyncScope element specifies whether just items or items and folder associated information are returned in a synchronization response.
	SyncScope *SyncScope `xml:"SyncScope"`
	// The SyncState element contains a base64-encoded form of the synchronization data that is updated after each successful request. This is used to identify the synchronization state.
	SyncState *SyncState `xml:"SyncState"`
}
