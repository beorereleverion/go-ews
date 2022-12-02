package elements

// The Companies element represents the collection of companies that are associated with a contact or task.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/companies
type Companies struct {
	// The String element represents a string that is used by items, contacts, tasks, and conversations.
	String *String `xml:"t:String"`
}
