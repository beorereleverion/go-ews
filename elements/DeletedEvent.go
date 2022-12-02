package elements

// The DeletedEvent element represents an event in which an item or folder is deleted.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/deletedevent
type DeletedEvent struct {
	// The FolderId element contains the identifier and change key of a folder.
	FolderId *FolderId `xml:"t:FolderId"`
	// The ItemId element contains the unique identifier and change key of an item in the Exchange store.
	ItemId *ItemId `xml:"t:ItemId"`
	// The ParentFolderId element represents the identifier of the parent folder that contains the item or folder.
	ParentFolderId *ParentFolderId `xml:"t:ParentFolderId"`
	// The Timestamp element represents the timestamp of a mailbox event.
	TimeStamp *TimeStamp `xml:"t:TimeStamp"`
	// The Watermark element represents an event bookmark in the mailbox event queue.
	Watermark *Watermark `xml:"t:Watermark"`
}
