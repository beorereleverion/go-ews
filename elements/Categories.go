package elements

// The Categories element contains a collection of strings that identify the categories to which an item in the mailbox belongs.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/categories-ex15websvcsotherref
type Categories struct {
	// The String element represents a string that is used by items, contacts, tasks, and conversations.
	String *String `xml:"t:String"`
}
