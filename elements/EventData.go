package elements

// The EventData element represents data that is associated with the processing step for the event.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/eventdata
type EventData struct {
	// The String element represents a string that is used by items, contacts, tasks, and conversations.
	String *String `xml:"t:String"`
}
