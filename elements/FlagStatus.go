package elements

// The FlagStatus element contains the aggregated flag status for conversation items in the current folder.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/flagstatus
type FlagStatus string

const (
	// Complete - Indicates the complete flag status.

	FlagStatusComplete FlagStatus = `Complete`
	// Flagged - Indicates the flagged status.

	FlagStatusFlagged FlagStatus = `Flagged`
	// NotFlagged - Indicates the not-flagged status.

	FlagStatusNotFlagged FlagStatus = `NotFlagged`
)
