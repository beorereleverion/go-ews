package elements

// The AssignCategories element represents the categories that are stamped on e-mail messages.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/assigncategories
type AssignCategories struct {
	// The String element represents a string that is used by items, contacts, tasks, and conversations.
	String *String `xml:"t:String"`
}
