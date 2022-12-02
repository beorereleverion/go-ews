package elements

// The UpdatedItemIds element specifies the identifiers of updated reminder items.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/updateditemids
type UpdatedItemIds struct {
	// The ItemId element contains the unique identifier and change key of an item in the Exchange store.
	ItemId *ItemId `xml:"t:ItemId"`
}
