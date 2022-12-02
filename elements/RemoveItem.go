package elements

// The RemoveItem element represents a response object that is used to remove a meeting item when a MeetingCancellation message is received.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/removeitem
type RemoveItem struct {
	// The ReferenceItemId element identifies the item to which the response object refers.
	ReferenceItemId *ReferenceItemId `xml:"t:ReferenceItemId"`
	// Represents the name of the RemoveItem reply object class as an English string.
	ObjectName *string `xml:"ObjectName,attr"`
}
