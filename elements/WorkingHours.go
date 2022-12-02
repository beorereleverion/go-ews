package elements

// The WorkingHours element represents the time zone settings and working hours for the requested mailbox user.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/workinghours-ex15websvcsotherref
type WorkingHours struct {
	// The TimeZone element contains elements that identify time zone information. This element also contains information about the transition between standard time and daylight saving time.
	TimeZone *TimeZoneAvailability `xml:"t:TimeZone"`
	// The WorkingPeriodArray element contains working period information for the mailbox user.
	WorkingPeriodArray *WorkingPeriodArray `xml:"t:WorkingPeriodArray"`
}
