package elements

// The ContainsBodyStrings element indicates the strings that must appear in the body of incoming messages in order for the condition or exception to apply.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/containsbodystrings
type ContainsBodyStrings struct {
	// The String element represents a string that is used by items, contacts, tasks, and conversations.
	String *String `xml:"t:String"`
}
