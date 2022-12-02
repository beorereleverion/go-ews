package elements

// The RequestedExtensionIds element contains an array of extension identifiers.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/requestedextensionids
type RequestedExtensionIds struct {
	// The String element represents a string that is used by items, contacts, tasks, and conversations.
	String *String `xml:"t:String"`
}
