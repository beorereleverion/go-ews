package elements

// The DraftItemIds element contains an array of item identifiers to draft items in a conversation.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/draftitemids
type DraftItemIds struct {
	// The ItemId element contains the unique identifier and change key of an item in the Exchange store.
	ItemId *ItemId `xml:"t:ItemId"`
	// The OccurrenceItemId element identifies a single occurrence of a recurring item.
	OccurrenceItemId *OccurrenceItemId `xml:"t:OccurrenceItemId"`
	// The RecurringMasterItemId element identifies a recurrence master item by identifying the identifiers of one of its related occurrence items.
	RecurringMasterItemId *RecurringMasterItemId `xml:"t:RecurringMasterItemId"`
	// The RecurringMasterItemIdRanges element specifies an array of occurrence ranges.
	RecurringMasterItemIdRanges *RecurringMasterItemIdRanges `xml:"t:RecurringMasterItemIdRanges"`
}
