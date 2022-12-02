package elements

// The WeeklyRecurrence element describes a weekly recurrence pattern.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/weeklyrecurrence
type WeeklyRecurrence struct {
	// The DaysOfWeek (DaysOfWeekType) element describes days of the week that are used in item recurrence patterns.
	DaysOfWeek *DaysOfWeekDaysOfWeekType `xml:"t:DaysOfWeek"`
	// The FirstDayOfWeek element specifies the first day of the week.
	FirstDayOfWeek *FirstDayOfWeek `xml:"t:FirstDayOfWeek"`
	// The Interval element defines the interval between two consecutive recurring items.
	Interval *Interval `xml:"t:Interval"`
}
