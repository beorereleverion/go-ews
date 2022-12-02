package elements

// The HasStartTimeChanged element specifies whether the start time for a meeting has changed.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/hasstarttimechanged
type HasStartTimeChanged bool

const (
	// false
	HasStartTimeChangedfalse HasStartTimeChanged = false
	// true
	HasStartTimeChangedtrue HasStartTimeChanged = true
)
