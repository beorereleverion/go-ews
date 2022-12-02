package elements

// The GlobalCategories element contains the category list for all conversation items in a mailbox.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/globalcategories
type GlobalCategories struct {
	// The String element represents a string that is used by items, contacts, tasks, and conversations.
	String *String `xml:"t:String"`
}
