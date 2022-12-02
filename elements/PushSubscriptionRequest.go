package elements

// The PushSubscriptionRequest element represents a subscription to a push-based event notification subscription.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/pushsubscriptionrequest
type PushSubscriptionRequest struct {
	// The EventTypes element contains a collection of event notification types that are used to create a subscription.
	EventTypes *EventTypes `xml:"t:EventTypes"`
	// The FolderIds element contains an array of folder identifiers that are used to identify folders to copy, move, get, delete, or monitor for event notifications.
	FolderIds *FolderIds `xml:"m:FolderIds"`
	// The StatusFrequency element represents the maximum timeout value, in minutes, in which retries are attempted by the server.
	StatusFrequency *StatusFrequency `xml:"t:StatusFrequency"`
	// The Url element represents the location of the client Web service for push notifications.
	Url *Url `xml:"Url"`
	// The Watermark element represents an event bookmark in the mailbox event queue.
	Watermark *Watermark `xml:"t:Watermark"`
	// Indicates whether to subscribe to all available folders. This attribute is optional.
	SubscribeToAllFolders *string `xml:"SubscribeToAllFolders,attr"`
}
