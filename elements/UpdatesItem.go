package elements

// The Updates element contains a set of elements that define append, set, and delete changes to item properties.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/updates-item
import "encoding/xml"

type UpdatesItem struct {
	XMLName xml.Name

	// The AppendToItemField element identifies data to append to a single property of an item during an UpdateItem operation.
	AppendToItemField *AppendToItemField `xml:"AppendToItemField"`
	// The DeleteItemField element represents an operation to delete a given property from an item during an UpdateItem call.
	DeleteItemField *DeleteItemField `xml:"DeleteItemField"`
	// The SetItemField element represents an update to a single property of an item in an UpdateItem operation.
	SetItemField *SetItemField `xml:"SetItemField"`
}
