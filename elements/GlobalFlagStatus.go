package elements

// The GlobalFlagStatus element contains the aggregated flag status for all conversation items in a mailbox.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/globalflagstatus
type GlobalFlagStatus string

const (
	// Complete - Indicates the complete flag status.

	GlobalFlagStatusComplete GlobalFlagStatus = `Complete`
	// Flagged - Indicates the flagged status.

	GlobalFlagStatusFlagged GlobalFlagStatus = `Flagged`
	// NotFlagged - Indicates the not-flagged status.

	GlobalFlagStatusNotFlagged GlobalFlagStatus = `NotFlagged`
)
