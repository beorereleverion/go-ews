package elements

// The IsWorkTime element represents whether the suggested meeting time occurs during scheduled work hours.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/isworktime
type IsWorkTime bool

const (
	// false
	IsWorkTimefalse IsWorkTime = false
	// true
	IsWorkTimetrue IsWorkTime = true
)
