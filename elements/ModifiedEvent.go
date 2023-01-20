package elements

// The ModifiedEvent element represents an event in which an item or folder is modified.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/modifiedevent
import "encoding/xml"

type ModifiedEvent struct {
	XMLName xml.Name

	// The FolderId element contains the identifier and change key of a folder.
	FolderId *FolderId `xml:"FolderId"`
	// The ItemId element contains the unique identifier and change key of an item in the Exchange store.
	ItemId *ItemId `xml:"ItemId"`
	// The ParentFolderId element represents the identifier of the parent folder that contains the item or folder.
	ParentFolderId *ParentFolderId `xml:"ParentFolderId"`
	// The Timestamp element represents the timestamp of a mailbox event.
	TimeStamp *TimeStamp `xml:"TimeStamp"`
	// The UnreadCount element contains the count of unread items within a folder.
	UnreadCount *UnreadCount `xml:"UnreadCount"`
	// The Watermark element represents an event bookmark in the mailbox event queue.
	Watermark *Watermark `xml:"Watermark"`
}

func (M *ModifiedEvent) SetForMarshal() {
	M.XMLName.Local = "t:ModifiedEvent"
}

func (M *ModifiedEvent) GetSchema() *Schema {
	return &SchemaTypes
}
