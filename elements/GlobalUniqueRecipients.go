package elements

// The GlobalUniqueRecipients element contains the recipient list of a conversation aggregated across a mailbox.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/globaluniquerecipients
type GlobalUniqueRecipients struct {
	// The String element represents a string that is used by items, contacts, tasks, and conversations.
	String *String `xml:"t:String"`
}
