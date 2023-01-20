package elements

// The Delete element identifies a single item to delete in the local client store.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/delete-itemsync
import "encoding/xml"

type DeleteItemSync struct {
	XMLName xml.Name

	// The ItemId element contains the unique identifier and change key of an item in the Exchange store.
	ItemId *ItemId `xml:"ItemId"`
}

func (D *DeleteItemSync) SetForMarshal() {
	D.XMLName.Local = "t:Delete"
}

func (D *DeleteItemSync) GetSchema() *Schema {
	return &SchemaTypes
}
