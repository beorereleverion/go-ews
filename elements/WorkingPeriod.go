package elements

// The WorkingPeriod element contains the work week days and hours of the mailbox user.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/workingperiod
type WorkingPeriod struct {
	// The DayOfWeek element contains the list of working days scheduled for the mailbox user.
	DayOfWeek *DayOfWeekWorkingPeriod `xml:"t:DayOfWeek"`
	// The EndTimeInMinutes element represents the end of the working day for a mailbox user.
	EndTimeInMinutes *EndTimeInMinutes `xml:"t:EndTimeInMinutes"`
	// The StartTimeInMinutes element represents the start of the working day for a mailbox user.
	StartTimeInMinutes *StartTimeInMinutes `xml:"t:StartTimeInMinutes"`
}
