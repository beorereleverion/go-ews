package elements

// The EditItems element indicates which items in a folder a user has permission to edit. This element was introduced in Microsoft Exchange Server 2007 Service Pack 1 (SP1).
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/edititems
import "encoding/xml"

type EditItems struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

const (
	// All
	EditItemsAll string = `All`
	// None
	EditItemsNone string = `None`
	// Owned
	EditItemsOwned string = `Owned`
)

func (E *EditItems) SetForMarshal() {
	E.XMLName.Local = "t:EditItems"
}

func (E *EditItems) GetSchema() *Schema {
	return &SchemaTypes
}
