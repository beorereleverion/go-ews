package elements

// The UpdatedItemIds element specifies the identifiers of updated reminder items.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/updateditemids
import "encoding/xml"

type UpdatedItemIds struct {
	XMLName xml.Name

	// The ItemId element contains the unique identifier and change key of an item in the Exchange store.
	ItemId *ItemId `xml:"ItemId"`
}

func (U *UpdatedItemIds) SetForMarshal() {
	U.XMLName.Local = "m:UpdatedItemIds"
}

func (U *UpdatedItemIds) GetSchema() *Schema {
	return &SchemaMessages
}
