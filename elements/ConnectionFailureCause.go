package elements

// The ConnectionFailureCause element specifies the reason for a disconnection from a telephone call.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/connectionfailurecause
type ConnectionFailureCause string

const (
	// NoAnswer
	ConnectionFailureCauseNoAnswer ConnectionFailureCause = `NoAnswer`
	// None
	ConnectionFailureCauseNone ConnectionFailureCause = `None`
	// Other
	ConnectionFailureCauseOther ConnectionFailureCause = `Other`
	// Unavailable
	ConnectionFailureCauseUnavailable ConnectionFailureCause = `Unavailable`
	// UserBusy
	ConnectionFailureCauseUserBusy ConnectionFailureCause = `UserBusy`
)
