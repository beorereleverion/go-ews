package elements

// The DaylightTime element represents an offset from the time relative to Coordinated Universal Time (UTC) that is represented by the Bias (UTC) element in regions where daylight saving time is observed. This element also contains information about when the transition to daylight saving time from standard time occurs.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/daylighttime
type DaylightTime struct {
	// The Bias element represents the offset from the Coordinated Universal Time (UTC) offset identified by the Bias (UTC) element for standard time and daylight saving time. This value is in minutes.
	Bias *Bias `xml:"t:Bias"`
	// The DayOfWeek element represents the day of the week on which the time zone transition occurs.
	DayOfWeek *DayOfWeekTimeZone `xml:"t:DayOfWeek"`
	// The DayOrder element represents the nth occurrence of the day specified in the DayOfWeek (TimeZone) element that represents the date of transition from and to standard time and daylight saving time.
	DayOrder *DayOrder `xml:"t:DayOrder"`
	// The Month element represents the transition month of the year to and from standard time and daylight saving time.
	Month *Month `xml:"t:Month"`
	// The Time element represents the transition time of day to and from standard time and daylight saving time.
	Time *Time `xml:"t:Time"`
	// The Year element is used to define a time zone that changes depending on the year. This element is optional. This element was introduced in Microsoft Exchange Server 2007 Service Pack 1 (SP1).
	Year *Year `xml:"t:Year"`
}
