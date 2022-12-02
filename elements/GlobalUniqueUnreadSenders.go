package elements

// The GlobalUniqueUnreadSenders element specifies a list of all the people who have sent messages that are currently unread in this conversation across all folders in the mailbox.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/globaluniqueunreadsenders
type GlobalUniqueUnreadSenders struct {
	// The String element represents a string that is used by items, contacts, tasks, and conversations.
	String *String `xml:"t:String"`
}
