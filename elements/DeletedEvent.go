package elements

// The DeletedEvent element represents an event in which an item or folder is deleted.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/deletedevent
import "encoding/xml"

type DeletedEvent struct {
	XMLName xml.Name

	// The FolderId element contains the identifier and change key of a folder.
	FolderId *FolderId `xml:"FolderId"`
	// The ItemId element contains the unique identifier and change key of an item in the Exchange store.
	ItemId *ItemId `xml:"ItemId"`
	// The ParentFolderId element represents the identifier of the parent folder that contains the item or folder.
	ParentFolderId *ParentFolderId `xml:"ParentFolderId"`
	// The Timestamp element represents the timestamp of a mailbox event.
	TimeStamp *TimeStamp `xml:"TimeStamp"`
	// The Watermark element represents an event bookmark in the mailbox event queue.
	Watermark *Watermark `xml:"Watermark"`
}

func (D *DeletedEvent) SetForMarshal() {
	D.XMLName.Local = "t:DeletedEvent"
}

func (D *DeletedEvent) GetSchema() *Schema {
	return &SchemaTypes
}
