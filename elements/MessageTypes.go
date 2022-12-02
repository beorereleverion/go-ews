package elements

// The MessageTypes element specifies an array of messages to search.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/messagetypes
type MessageTypes struct {
	// The SearchItemKind element indicates the type of items that are searched for a FindMailboxStatisticsByKeyword operation.
	SearchItemKind *SearchItemKind `xml:"t:SearchItemKind"`
}
