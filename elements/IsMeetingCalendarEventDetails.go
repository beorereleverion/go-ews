package elements

// The IsMeeting element indicates whether the calendar event is a meeting or an appointment.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/ismeeting-calendareventdetails
import "encoding/xml"

type IsMeetingCalendarEventDetails struct {
	XMLName xml.Name
	TEXT    bool `xml:",chardata"`
}

const (
	// false
	IsMeetingCalendarEventDetailsfalse bool = false
	// true
	IsMeetingCalendarEventDetailstrue bool = true
)

func (I *IsMeetingCalendarEventDetails) SetForMarshal() {
	I.XMLName.Local = "t:IsMeeting"
}

func (I *IsMeetingCalendarEventDetails) GetSchema() *Schema {
	return &SchemaTypes
}
