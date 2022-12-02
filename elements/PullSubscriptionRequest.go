package elements

// The PullSubscriptionRequest element represents a subscription to a pull-based event notification subscription.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/pullsubscriptionrequest
type PullSubscriptionRequest struct {
	// The EventTypes element contains a collection of event notification types that are used to create a subscription.
	EventTypes *EventTypes `xml:"t:EventTypes"`
	// The FolderIds element contains an array of folder identifiers that are used to identify folders to copy, move, get, delete, or monitor for event notifications.
	FolderIds *FolderIds `xml:"m:FolderIds"`
	// The Timeout element represents the duration, in minutes, that the subscription can remain idle without a GetEvents request from the client.
	Timeout *Timeout `xml:"t:Timeout"`
	// The Watermark element represents an event bookmark in the mailbox event queue.
	Watermark *Watermark `xml:"t:Watermark"`
	// Indicates whether to subscribe to all available folders. This attribute is optional.
	SubscribeToAllFolders *string `xml:"SubscribeToAllFolders,attr"`
}
