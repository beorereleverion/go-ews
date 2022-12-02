package elements

// The CalendarEventDetails element provides additional information about a calendar event.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/calendareventdetails
type CalendarEventDetails struct {
	// The ID element represents the entry ID of the calendar item.
	ID *ID `xml:"t:ID"`
	// The IsException element indicates whether an instance of a recurring calendar item is changed from the master.
	IsException *IsException `xml:"t:IsException"`
	// The IsMeeting element indicates whether the calendar event is a meeting or an appointment.
	IsMeeting *IsMeetingCalendarEventDetails `xml:"t:IsMeeting"`
	// The IsPrivate element indicates whether the calendar item is private.
	IsPrivate *IsPrivate `xml:"t:IsPrivate"`
	// The IsRecurring element indicates whether the calendar event is an instance of a recurring calendar item or a single calendar item.
	IsRecurring *IsRecurringCalendarEventDetails `xml:"t:IsRecurring"`
	// The IsReminderSet element indicates whether a reminder has been set for the calendar event.
	IsReminderSet *IsReminderSet `xml:"t:IsReminderSet"`
	// The Location element represents the location field of the calendar item.
	Location *LocationCalendarEventDetails `xml:"t:Location"`
	// The Subject element represents the subject of a calendar item.
	Subject *SubjectCalendarEventDetails `xml:"t:Subject"`
}
