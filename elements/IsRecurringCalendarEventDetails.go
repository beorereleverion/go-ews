package elements

// The IsRecurring element indicates whether the calendar event is an instance of a recurring calendar item or a single calendar item.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/isrecurring-calendareventdetails
type IsRecurringCalendarEventDetails bool

const (
	// false
	IsRecurringCalendarEventDetailsfalse IsRecurringCalendarEventDetails = false
	// true
	IsRecurringCalendarEventDetailstrue IsRecurringCalendarEventDetails = true
)
