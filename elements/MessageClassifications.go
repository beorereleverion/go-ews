package elements

// The MessageClassifications element represents the message classifications that must be stamped on incoming messages in order for the condition or exception to apply.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/messageclassifications
type MessageClassifications struct {
	// The String element represents a string that is used by items, contacts, tasks, and conversations.
	String *String `xml:"t:String"`
}
