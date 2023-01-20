package elements

// The CalendarEventDetails element provides additional information about a calendar event.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/calendareventdetails
import "encoding/xml"

type CalendarEventDetails struct {
	XMLName xml.Name

	// The ID element represents the entry ID of the calendar item.
	ID *ID `xml:"ID"`
	// The IsException element indicates whether an instance of a recurring calendar item is changed from the master.
	IsException *IsException `xml:"IsException"`
	// The IsMeeting element indicates whether the calendar event is a meeting or an appointment.
	IsMeeting *IsMeetingCalendarEventDetails `xml:"IsMeeting"`
	// The IsPrivate element indicates whether the calendar item is private.
	IsPrivate *IsPrivate `xml:"IsPrivate"`
	// The IsRecurring element indicates whether the calendar event is an instance of a recurring calendar item or a single calendar item.
	IsRecurring *IsRecurringCalendarEventDetails `xml:"IsRecurring"`
	// The IsReminderSet element indicates whether a reminder has been set for the calendar event.
	IsReminderSet *IsReminderSet `xml:"IsReminderSet"`
	// The Location element represents the location field of the calendar item.
	Location *LocationCalendarEventDetails `xml:"Location"`
	// The Subject element represents the subject of a calendar item.
	Subject *SubjectCalendarEventDetails `xml:"Subject"`
}

func (C *CalendarEventDetails) SetForMarshal() {
	C.XMLName.Local = "t:CalendarEventDetails"
}

func (C *CalendarEventDetails) GetSchema() *Schema {
	return &SchemaTypes
}
