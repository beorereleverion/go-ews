package elements

// The RecipientTrackingEvents element represents a collection of one or more events for a message.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/recipienttrackingevents
type RecipientTrackingEvents struct {
	// The RecipientTrackingEvent element contains information for a single event for a recipient.
	RecipientTrackingEvent *RecipientTrackingEvent `xml:"t:RecipientTrackingEvent"`
}
