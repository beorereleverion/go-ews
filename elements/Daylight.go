package elements

// The Daylight element represents the date and time when the time changes from standard time to daylight saving time.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/daylight
type Daylight struct {
	// The AbsoluteDate element represents the date when the time changes from standard or daylight saving time.
	AbsoluteDate *AbsoluteDate `xml:"t:AbsoluteDate"`
	// The BaseOffset element represents the hourly offset from Coordinated Universal Time (UTC) for the current time zone.
	BaseOffset *BaseOffset `xml:"t:BaseOffset"`
	// The Offset element describes the offset from the BaseOffset. Together with the BaseOffset element, the Offset element identifies whether the time is standard or daylight saving time.
	Offset *Offset `xml:"t:Offset"`
	// The RelativeYearlyRecurrence element describes a relative yearly recurrence pattern.
	RelativeYearlyRecurrence *RelativeYearlyRecurrence `xml:"t:RelativeYearlyRecurrence"`
	// The Time element describes the time when the time changes between standard time and daylight saving time.
	Time *TimeTimeChangeType `xml:"t:Time"`
	// Describes the name of the time zone.
	TimeZoneName *string `xml:"TimeZoneName,attr"`
}
