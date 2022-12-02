package elements

// The SortOrder element defines how items are sorted in a FindItem or FindConversation request.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/sortorder
type SortOrder struct {
	// The FieldOrder element represents a single field by which to sort results and indicates the direction for the sort.
	FieldOrder *FieldOrder `xml:"t:FieldOrder"`
}
