package elements

// The StreamingSubscriptionRequest element represents a subscription to a streaming event notification subscription.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/streamingsubscriptionrequest
import "encoding/xml"

type StreamingSubscriptionRequest struct {
	XMLName xml.Name

	// The EventTypes element contains a collection of event notification types that are used to create a subscription.
	EventTypes *EventTypes `xml:"EventTypes"`
	// The FolderIds element contains an array of folder identifiers that are used to identify folders to copy, move, get, delete, or monitor for event notifications.
	FolderIds *FolderIds `xml:"FolderIds"`
	// Indicates whether the server will subscribe to all folders in the user's mailbox. A value of true indicates that the server will subscribe.
	SubscribeToAllFolders *string `xml:"SubscribeToAllFolders,attr"`
}

func (S *StreamingSubscriptionRequest) SetForMarshal() {
	S.XMLName.Local = "m:StreamingSubscriptionRequest"
}

func (S *StreamingSubscriptionRequest) GetSchema() *Schema {
	return &SchemaMessages
}
