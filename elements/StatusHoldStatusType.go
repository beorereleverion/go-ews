package elements

// The Status element specifies the hold status for a mailbox.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/status-holdstatustype
type StatusHoldStatusType string

const (
	// Failed
	StatusHoldStatusTypeFailed StatusHoldStatusType = `Failed`
	// NotOnHold
	StatusHoldStatusTypeNotOnHold StatusHoldStatusType = `NotOnHold`
	// OnHold
	StatusHoldStatusTypeOnHold StatusHoldStatusType = `OnHold`
	// PartialHold
	StatusHoldStatusTypePartialHold StatusHoldStatusType = `PartialHold`
	// Pending
	StatusHoldStatusTypePending StatusHoldStatusType = `Pending`
)
