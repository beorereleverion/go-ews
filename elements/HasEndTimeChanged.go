package elements

// The HasEndTimeChanged element specifies whether the end time for a meeting has changed.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/hasendtimechanged
type HasEndTimeChanged bool

const (
	// false
	HasEndTimeChangedfalse HasEndTimeChanged = false
	// true
	HasEndTimeChangedtrue HasEndTimeChanged = true
)
