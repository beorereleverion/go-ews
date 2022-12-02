package elements

// The IsMeeting element indicates whether the calendar event is a meeting or an appointment.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/ismeeting-calendareventdetails
type IsMeetingCalendarEventDetails bool

const (
	// false
	IsMeetingCalendarEventDetailsfalse IsMeetingCalendarEventDetails = false
	// true
	IsMeetingCalendarEventDetailstrue IsMeetingCalendarEventDetails = true
)
