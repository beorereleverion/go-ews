package elements

// The NewMailEvent element represents an event that is triggered by a new mail item in a mailbox.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/newmailevent
import "encoding/xml"

type NewMailEvent struct {
	XMLName xml.Name

	// The ItemId element contains the unique identifier and change key of an item in the Exchange store.
	ItemId *ItemId `xml:"ItemId"`
	// The ParentFolderId element represents the identifier of the parent folder that contains the item or folder.
	ParentFolderId *ParentFolderId `xml:"ParentFolderId"`
	// The Timestamp element represents the timestamp of a mailbox event.
	TimeStamp *TimeStamp `xml:"TimeStamp"`
	// The Watermark element represents an event bookmark in the mailbox event queue.
	Watermark *Watermark `xml:"Watermark"`
}

func (N *NewMailEvent) SetForMarshal() {
	N.XMLName.Local = "t:NewMailEvent"
}

func (N *NewMailEvent) GetSchema() *Schema {
	return &SchemaTypes
}
