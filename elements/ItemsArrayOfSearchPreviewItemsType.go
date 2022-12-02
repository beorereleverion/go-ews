package elements

// The Items element specifies a list of items available for preview as the results of a SearchMailboxes operation.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/items-arrayofsearchpreviewitemstype
type ItemsArrayOfSearchPreviewItemsType struct {
	// The SearchPreviewItem element specifies the item preview for a discovery search.
	SearchPreviewItem *SearchPreviewItem `xml:"t:SearchPreviewItem"`
}
