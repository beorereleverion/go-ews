package elements

// The AppointmentReplyTime element represents the date and time that an attendee replied to a meeting request.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/appointmentreplytime
import "time"

import "encoding/xml"

type AppointmentReplyTime struct {
	XMLName xml.Name
	TEXT    time.Time `xml:",chardata"`
}

func (A *AppointmentReplyTime) SetForMarshal() {
	A.XMLName.Local = "t:AppointmentReplyTime"
}

func (A *AppointmentReplyTime) GetSchema() *Schema {
	return &SchemaTypes
}
