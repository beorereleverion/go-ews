package elements

// The TimeZone element contains elements that identify time zone information. This element also contains information about the transition between standard time and daylight saving time.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/timezone-availability
type TimeZoneAvailability struct {
	// The DaylightTime element represents an offset from the time relative to Coordinated Universal Time (UTC) that is represented by the Bias (UTC) element in regions where daylight saving time is observed. This element also contains information about when the transition to daylight saving time from standard time occurs.
	DaylightTime *DaylightTime `xml:"t:DaylightTime"`
	// The StandardTime element represents an offset from the time relative to Coordinated Universal Time (UTC) that is represented by the Bias (UTC) element. This element also contains information about the transition to standard time from daylight saving time in regions where daylight saving time is observed.
	StandardTime *StandardTime `xml:"t:StandardTime"`
}
