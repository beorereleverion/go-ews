package elements

// The RootItemId element identifies the root item of a deleted attachment.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/rootitemid
type RootItemId struct {
	// Identifies the new change key of the root item of a deleted attachment.
	RootItemChangeKey *string `xml:"RootItemChangeKey,attr"`
	// Identifies the root item of a deleted attachment.
	RootItemId *string `xml:"RootItemId,attr"`
}
