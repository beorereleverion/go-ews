package elements

// The Children element contains the names of a contact's children.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/children
type Children struct {
	// The String element represents a string that is used by items, contacts, tasks, and conversations.
	String *String `xml:"t:String"`
}
