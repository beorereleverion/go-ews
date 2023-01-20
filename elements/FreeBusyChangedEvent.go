package elements

// The FreeBusyChangedEvent element represents an event in which an item's free/busy time has changed.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/freebusychangedevent
import "encoding/xml"

type FreeBusyChangedEvent struct {
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

func (F *FreeBusyChangedEvent) SetForMarshal() {
	F.XMLName.Local = "t:FreeBusyChangedEvent"
}

func (F *FreeBusyChangedEvent) GetSchema() *Schema {
	return &SchemaTypes
}
