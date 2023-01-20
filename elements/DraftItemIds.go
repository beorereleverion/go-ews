package elements

// The DraftItemIds element contains an array of item identifiers to draft items in a conversation.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/draftitemids
import "encoding/xml"

type DraftItemIds struct {
	XMLName xml.Name

	// The ItemId element contains the unique identifier and change key of an item in the Exchange store.
	ItemId *ItemId `xml:"ItemId"`
	// The OccurrenceItemId element identifies a single occurrence of a recurring item.
	OccurrenceItemId *OccurrenceItemId `xml:"OccurrenceItemId"`
	// The RecurringMasterItemId element identifies a recurrence master item by identifying the identifiers of one of its related occurrence items.
	RecurringMasterItemId *RecurringMasterItemId `xml:"RecurringMasterItemId"`
	// The RecurringMasterItemIdRanges element specifies an array of occurrence ranges.
	RecurringMasterItemIdRanges *RecurringMasterItemIdRanges `xml:"RecurringMasterItemIdRanges"`
}

func (D *DraftItemIds) SetForMarshal() {
	D.XMLName.Local = "t:DraftItemIds"
}

func (D *DraftItemIds) GetSchema() *Schema {
	return &SchemaTypes
}
