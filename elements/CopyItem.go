package elements

// The CopyItem element defines a request to copy an item in a mailbox in the Exchange store.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/copyitem
import "encoding/xml"

type CopyItem struct {
	XMLName xml.Name

	// The ItemIds element contains the unique identities of items, occurrence items, and recurring master items that are used to delete, send, get, move, or copy items in the Exchange store.
	ItemIds *ItemIds `xml:"ItemIds"`
	// The ReturnNewItemIds element indicates whether the item identifiers of new items are returned in the response.
	ReturnNewItemIds *ReturnNewItemIds `xml:"ReturnNewItemIds"`
	// The ToFolderId element represents the destination folder for a copied or moved item or folder.
	ToFolderId *ToFolderId `xml:"ToFolderId"`
}

func (C *CopyItem) SetForMarshal() {
	C.XMLName.Local = "m:CopyItem"
}

func (C *CopyItem) GetSchema() *Schema {
	return &SchemaMessages
}
