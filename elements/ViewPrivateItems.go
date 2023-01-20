package elements

// The ViewPrivateItems element indicates whether a delegate user or client application has permission to view private items in the principal's mailbox.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/viewprivateitems
import "encoding/xml"

type ViewPrivateItems struct {
	XMLName xml.Name
	TEXT    bool `xml:",chardata"`
}

const (
	// false
	ViewPrivateItemsfalse bool = false
	// true
	ViewPrivateItemstrue bool = true
)

func (V *ViewPrivateItems) SetForMarshal() {
	V.XMLName.Local = "t:ViewPrivateItems"
}

func (V *ViewPrivateItems) GetSchema() *Schema {
	return &SchemaTypes
}
