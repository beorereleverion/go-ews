package elements

// The Contacts element contains a list of contacts that are associated with a task.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/contacts-ex15websvcsotherref
type Contacts struct {
	// The String element represents a string that is used by items, contacts, tasks, and conversations.
	String *String `xml:"t:String"`
}
