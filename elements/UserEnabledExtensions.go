package elements

// The UserEnabledExtensions element lists the enabled apps.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/userenabledextensions
type UserEnabledExtensions struct {
	// The String element represents a string that is used by items, contacts, tasks, and conversations.
	String *String `xml:"t:String"`
}
