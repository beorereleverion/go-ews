package elements

// The GlobalItemIds element contains the collection of item identifiers for all conversation items in a mailbox.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/globalitemids
import "encoding/xml"

type GlobalItemIds struct {
	XMLName xml.Name

	// The ItemId element contains the unique identifier and change key of an item in the Exchange store.
	ItemId *ItemId `xml:"ItemId"`
	// The OccurrenceItemId element identifies a single occurrence of a recurring item.
	OccurrenceItemId *OccurrenceItemId `xml:"OccurrenceItemId"`
	// The RecurringMasterItemId element identifies a recurrence master item by identifying the identifiers of one of its related occurrence items.
	RecurringMasterItemId *RecurringMasterItemId `xml:"RecurringMasterItemId"`
}

func (G *GlobalItemIds) SetForMarshal() {
	G.XMLName.Local = "t:GlobalItemIds"
}

func (G *GlobalItemIds) GetSchema() *Schema {
	return &SchemaTypes
}
