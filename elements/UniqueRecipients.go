package elements

// The UniqueRecipients element contains the recipient list of a conversation in a particular folder. This element is read-only.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/uniquerecipients
type UniqueRecipients struct {
	// The String element represents a string that is used by items, contacts, tasks, and conversations.
	String *String `xml:"t:String"`
}
