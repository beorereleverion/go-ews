package elements

// The ItemIds element contains an array of item identifiers that identify the items to export from a mailbox.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/itemids-nonemptyarrayofitemidstype
type ItemIdsNonEmptyArrayOfItemIdsType struct {
	// The ItemId element contains the unique identifier and change key of an item in the Exchange store.
	ItemId *ItemId `xml:"t:ItemId"`
}
