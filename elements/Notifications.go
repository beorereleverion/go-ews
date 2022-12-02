package elements

// The Notifications element contains an array of information about the subscription and the events that have occurred since the last notification.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/notifications
type Notifications struct {
	// The Notification element contains information about the subscription and the events that have occurred since the last notification.
	Notification *Notification `xml:"t:Notification"`
}
