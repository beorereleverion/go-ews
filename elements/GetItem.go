package elements

// The GetItem element defines a request to get an item from a mailbox in the Exchange store.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/getitem
import "encoding/xml"

type GetItem struct {
	XMLName xml.Name

	// The ItemIds element contains the unique identities of items, occurrence items, and recurring master items that are used to delete, send, get, move, or copy items in the Exchange store.
	ItemIds *ItemIds `xml:"ItemIds"`
	// The ItemShape element identifies a set of properties to return in a GetItem operation, FindItem operation, or SyncFolderItems operation response.
	ItemShape *ItemShape `xml:"ItemShape"`
}

func (G *GetItem) SetForMarshal() {
	G.XMLName.Local = "m:GetItem"
}

func (G *GetItem) GetSchema() *Schema {
	return &SchemaMessages
}
