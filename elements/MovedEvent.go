package elements

// The MovedEvent element represents an event in which an item or folder is moved from one parent folder to another parent folder.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/movedevent
type MovedEvent struct {
	// The FolderId element contains the identifier and change key of a folder.
	FolderId *FolderId `xml:"t:FolderId"`
	// The ItemId element contains the unique identifier and change key of an item in the Exchange store.
	ItemId *ItemId `xml:"t:ItemId"`
	// The OldFolderId element contains the original identifier of a folder that was moved or copied.
	OldFolderId *OldFolderId `xml:"t:OldFolderId"`
	// The OldItemId element contains the unique identifier of the item that was copied or moved.
	OldItemId *OldItemId `xml:"t:OldItemId"`
	// The OldParentFolderId element contains the identifier of the parent folder of an item or folder that was copied or moved.
	OldParentFolderId *OldParentFolderId `xml:"t:OldParentFolderId"`
	// The ParentFolderId element represents the identifier of the parent folder that contains the item or folder.
	ParentFolderId *ParentFolderId `xml:"t:ParentFolderId"`
	// The Timestamp element represents the timestamp of a mailbox event.
	TimeStamp *TimeStamp `xml:"t:TimeStamp"`
	// The Watermark element represents an event bookmark in the mailbox event queue.
	Watermark *Watermark `xml:"t:Watermark"`
}
