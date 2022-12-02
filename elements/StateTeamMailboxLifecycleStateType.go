package elements

// The State element contains the lifecycle state that is set on a site mailbox.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/state-teammailboxlifecyclestatetype
type StateTeamMailboxLifecycleStateType string

const (
	// Active
	StateTeamMailboxLifecycleStateTypeActive StateTeamMailboxLifecycleStateType = `Active`
	// Closed
	StateTeamMailboxLifecycleStateTypeClosed StateTeamMailboxLifecycleStateType = `Closed`
	// PendingDelete
	StateTeamMailboxLifecycleStateTypePendingDelete StateTeamMailboxLifecycleStateType = `PendingDelete`
	// Unlinked
	StateTeamMailboxLifecycleStateTypeUnlinked StateTeamMailboxLifecycleStateType = `Unlinked`
)
