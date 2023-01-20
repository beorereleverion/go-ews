package elements

// The DeleteItems element indicates which items in a folder a user has permission to delete. This element was introduced in Microsoft Exchange Server 2007 Service Pack 1 (SP1).
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/deleteitems
import "encoding/xml"

type DeleteItems struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

const (
	// All
	DeleteItemsAll string = `All`
	// None
	DeleteItemsNone string = `None`
	// Owned
	DeleteItemsOwned string = `Owned`
)

func (D *DeleteItems) SetForMarshal() {
	D.XMLName.Local = "t:DeleteItems"
}

func (D *DeleteItems) GetSchema() *Schema {
	return &SchemaTypes
}
