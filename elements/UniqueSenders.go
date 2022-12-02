package elements

// The UniqueSenders element contains a list of all the senders of conversation items in the current folder. This element is read-only.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/uniquesenders
type UniqueSenders struct {
	// The String element represents a string that is used by items, contacts, tasks, and conversations.
	String *String `xml:"t:String"`
}
