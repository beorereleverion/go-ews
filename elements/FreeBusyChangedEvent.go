package elements

// The FreeBusyChangedEvent element represents an event in which an item's free/busy time has changed.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/freebusychangedevent
type FreeBusyChangedEvent struct {
	// The ItemId element contains the unique identifier and change key of an item in the Exchange store.
	ItemId *ItemId `xml:"t:ItemId"`
	// The ParentFolderId element represents the identifier of the parent folder that contains the item or folder.
	ParentFolderId *ParentFolderId `xml:"t:ParentFolderId"`
	// The Timestamp element represents the timestamp of a mailbox event.
	TimeStamp *TimeStamp `xml:"t:TimeStamp"`
	// The Watermark element represents an event bookmark in the mailbox event queue.
	Watermark *Watermark `xml:"t:Watermark"`
}
