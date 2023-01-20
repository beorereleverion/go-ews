package elements

// The MarkAllItemsAsRead element contains the request to mark all the items in a folder as read.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/markallitemsasread
import "encoding/xml"

type MarkAllItemsAsRead struct {
	XMLName xml.Name

	// The FolderIds element contains an array of folder identifiers that are used to identify folders to copy, move, get, delete, or monitor for event notifications.
	FolderIds *FolderIds `xml:"FolderIds"`
	// The ReadFlag element indicates the read state to set on items in a folder.
	ReadFlag *ReadFlag `xml:"ReadFlag"`
	// The SuppressReadReceipts element indicates whether read receipts should be suppressed.
	SuppressReadReceipts *SuppressReadReceipts `xml:"SuppressReadReceipts"`
}

func (M *MarkAllItemsAsRead) SetForMarshal() {
	M.XMLName.Local = "m:MarkAllItemsAsRead"
}

func (M *MarkAllItemsAsRead) GetSchema() *Schema {
	return &SchemaMessages
}
