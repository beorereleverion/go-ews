package elements

// The UniqueUnreadSenders element contains a list of all the people who have sent messages that are currently unread in this conversation in the current folder. This element is read-only.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/uniqueunreadsenders
type UniqueUnreadSenders struct {
	// The String element represents a string that is used by items, contacts, tasks, and conversations.
	String *String `xml:"t:String"`
}
