package elements

// The ItemIds element contains the unique identities of items, occurrence items, and recurring master items that are used to delete, send, get, move, or copy items in the Exchange store.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/itemids
type ItemIds struct {
	// The ItemId element contains the unique identifier and change key of an item in the Exchange store.
	ItemId *ItemId `xml:"t:ItemId"`
	// The OccurrenceItemId element identifies a single occurrence of a recurring item.
	OccurrenceItemId *OccurrenceItemId `xml:"t:OccurrenceItemId"`
	// The RecurringMasterItemId element identifies a recurrence master item by identifying the identifiers of one of its related occurrence items.
	RecurringMasterItemId *RecurringMasterItemId `xml:"t:RecurringMasterItemId"`
}
