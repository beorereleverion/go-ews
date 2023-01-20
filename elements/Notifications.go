package elements

// The Notifications element contains an array of information about the subscription and the events that have occurred since the last notification.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/notifications
import "encoding/xml"

type Notifications struct {
	XMLName xml.Name

	// The Notification element contains information about the subscription and the events that have occurred since the last notification.
	Notification *Notification `xml:"Notification"`
}

func (N *Notifications) SetForMarshal() {
	N.XMLName.Local = "m:Notifications"
}

func (N *Notifications) GetSchema() *Schema {
	return &SchemaMessages
}
