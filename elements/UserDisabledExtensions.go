package elements

// The UserDisabledExtensions element lists the disabled apps.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/userdisabledextensions
type UserDisabledExtensions struct {
	// The String element represents a string that is used by items, contacts, tasks, and conversations.
	String *String `xml:"t:String"`
}
