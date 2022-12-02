package elements

// The Flag element specifies a flag on a mailbox item.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/flag
type Flag struct {
	// The CompleteDate element represents the date on which an item was completed.
	CompleteDate *CompleteDate `xml:"t:CompleteDate"`
	// The DueDate element represents the date an item is due.
	DueDate *DueDate `xml:"t:DueDate"`
	// The FlagStatus element contains the aggregated flag status for conversation items in the current folder.
	FlagStatus *FlagStatus `xml:"t:FlagStatus"`
	// The StartDate element represents the start date of an item.
	StartDate *StartDate `xml:"t:StartDate"`
}
