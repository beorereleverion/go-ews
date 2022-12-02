package elements

// The FreeBusyView element contains availability information for a specific user.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/freebusyview
type FreeBusyView struct {
	// The CalendarEventArray element contains a set of unique calendar item occurrences that represent the requested user's availability.
	CalendarEventArray *CalendarEventArray `xml:"t:CalendarEventArray"`
	// The FreeBusyViewType element represents the type of free/busy information returned in the response.
	FreeBusyViewType *FreeBusyViewType `xml:"t:FreeBusyViewType"`
	// The MergedFreeBusy element contains the merged free/busy stream of data.
	MergedFreeBusy *MergedFreeBusy `xml:"t:MergedFreeBusy"`
	// The WorkingHours element represents the time zone settings and working hours for the requested mailbox user.
	WorkingHours *WorkingHours `xml:"t:WorkingHours"`
}
