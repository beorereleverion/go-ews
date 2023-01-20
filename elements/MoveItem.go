package elements

// The MoveItem element defines a request to move an item in the Exchange store.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/moveitem
import "encoding/xml"

type MoveItem struct {
	XMLName xml.Name

	// The ItemIds element contains the unique identities of items, occurrence items, and recurring master items that are used to delete, send, get, move, or copy items in the Exchange store.
	ItemIds *ItemIds `xml:"ItemIds"`
	// The ReturnNewItemIds element indicates whether the item identifiers of new items are returned in the response.
	ReturnNewItemIds *ReturnNewItemIds `xml:"ReturnNewItemIds"`
	// The ToFolderId element represents the destination folder for a copied or moved item or folder.
	ToFolderId *ToFolderId `xml:"ToFolderId"`
}

func (M *MoveItem) SetForMarshal() {
	M.XMLName.Local = "m:MoveItem"
}

func (M *MoveItem) GetSchema() *Schema {
	return &SchemaMessages
}
