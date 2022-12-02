package elements

// The RecurringDateTransition element represents a time zone transition that occurs on a specific date each year.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/recurringdatetransition
type RecurringDateTransition struct {
	// The Day element represents the day of the month on which the time zone transition occurs.
	Day *Day `xml:"t:Day"`
	// The Month element represents the month in which the time zone transition occurs.
	Month *MonthTimeZoneTransition `xml:"t:Month"`
	// The Period element defines the name, time offset, and unique identifier for a specific stage of the time zone.
	Period *Period `xml:"t:Period"`
	// The TimeOffset element represents the time offset from Coordinated Universal Time (UTC) for the time zone transition.
	TimeOffset *TimeOffset `xml:"t:TimeOffset"`
	// The To element specifies the target of the time zone transition. The target is either a time zone period or a group of time zone transitions.
	To *To `xml:"t:To"`
	// The TransitionsGroup element represents an array of time zone transitions.
	TransitionsGroup *TransitionsGroup `xml:"t:TransitionsGroup"`
}
