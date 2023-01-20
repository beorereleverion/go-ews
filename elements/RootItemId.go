package elements

// The RootItemId element identifies the root item of a deleted attachment.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/rootitemid
import "encoding/xml"

type RootItemId struct {
	XMLName xml.Name

	// Identifies the new change key of the root item of a deleted attachment.
	RootItemChangeKey *string `xml:"RootItemChangeKey,attr"`
	// Identifies the root item of a deleted attachment.
	RootItemId *string `xml:"RootItemId,attr"`
}

func (R *RootItemId) SetForMarshal() {
	R.XMLName.Local = "t:RootItemId"
}

func (R *RootItemId) GetSchema() *Schema {
	return &SchemaTypes
}
