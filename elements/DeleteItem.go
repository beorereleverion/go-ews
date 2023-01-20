package elements

// The DeleteItem element defines a request to delete an item from a mailbox in the Exchange store.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/deleteitem
import "encoding/xml"

type DeleteItem struct {
	XMLName xml.Name

	// The ItemIds element contains the unique identities of items, occurrence items, and recurring master items that are used to delete, send, get, move, or copy items in the Exchange store.
	ItemIds *ItemIds `xml:"ItemIds"`
	// Describes whether a task instance or a task master is deleted by a DeleteItem operation. This attribute is required when tasks are deleted. This attribute is optional when non-task items are deleted.
	AffectedTaskOccurrences *string `xml:"AffectedTaskOccurrences,attr"`
	// Describes how an item is deleted. This attribute is required.
	DeleteType *string `xml:"DeleteType,attr"`
	// Describes whether a calendar item deletion is communicated to attendees. This attribute is required when calendar items are deleted. This attribute is optional if non-calendar items are deleted.
	SendMeetingCancellations *string `xml:"SendMeetingCancellations,attr"`
	// Indicates whether read receipts for the deleted item are suppressed. A text value of true, indicates that the read receipts are suppressed. A value of false indicates that the read receipts are sent to the sender. This attribute is optional.
	SuppressReadReceipts *string `xml:"SuppressReadReceipts,attr"`
}

func (D *DeleteItem) SetForMarshal() {
	D.XMLName.Local = "m:DeleteItem"
}

func (D *DeleteItem) GetSchema() *Schema {
	return &SchemaMessages
}
