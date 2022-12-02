package elements

// The MeetingTimeZone element represents the time zone of the location where the meeting is hosted.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/meetingtimezone
type MeetingTimeZone struct {
	// The BaseOffset element represents the hourly offset from Coordinated Universal Time (UTC) for the current time zone.
	BaseOffset *BaseOffset `xml:"t:BaseOffset"`
	// The Daylight element represents the date and time when the time changes from standard time to daylight saving time.
	Daylight *Daylight `xml:"t:Daylight"`
	// The Standard element represents the date and time when the time changes from daylight saving time to standard time.
	Standard *Standard `xml:"t:Standard"`
	// Describes the name of the time zone.
	TimeZoneName *string `xml:"TimeZoneName,attr"`
}
