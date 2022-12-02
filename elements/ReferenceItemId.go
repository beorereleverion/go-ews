package elements

// The ReferenceItemId element identifies the item to which the response object refers.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/referenceitemid
type ReferenceItemId struct {
	// Identifies a specific version of an item.
	ChangeKey *string `xml:"ChangeKey,attr"`
	// Identifies a specific item in the Exchange store.
	Id *string `xml:"Id,attr"`
}
