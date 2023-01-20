package elements

// The ItemChange element contains an item identifier and the updates to apply to the item.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/itemchange
import "encoding/xml"

type ItemChange struct {
	XMLName xml.Name

	// The ItemId element contains the unique identifier and change key of an item in the Exchange store.
	ItemId *ItemId `xml:"ItemId"`
	// The OccurrenceItemId element identifies a single occurrence of a recurring item.
	OccurrenceItemId *OccurrenceItemId `xml:"OccurrenceItemId"`
	// The RecurringMasterItemId element identifies a recurrence master item by identifying the identifiers of one of its related occurrence items.
	RecurringMasterItemId *RecurringMasterItemId `xml:"RecurringMasterItemId"`
	// The Updates element contains a set of elements that define append, set, and delete changes to item properties.
	Updates *UpdatesItem `xml:"Updates"`
}

func (I *ItemChange) SetForMarshal() {
	I.XMLName.Local = "t:ItemChange"
}

func (I *ItemChange) GetSchema() *Schema {
	return &SchemaTypes
}
