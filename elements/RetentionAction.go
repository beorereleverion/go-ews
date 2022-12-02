package elements

// The RetentionAction element specifies the action performed on items with the retention tag.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/retentionaction
type RetentionAction string

const (
	// DeleteAndAllowRecovery
	RetentionActionDeleteAndAllowRecovery RetentionAction = `DeleteAndAllowRecovery`
	// MarkAsPastRetentionLimit
	RetentionActionMarkAsPastRetentionLimit RetentionAction = `MarkAsPastRetentionLimit`
	// MoveToArchive
	RetentionActionMoveToArchive RetentionAction = `MoveToArchive`
	// MoveToDeletedItems
	RetentionActionMoveToDeletedItems RetentionAction = `MoveToDeletedItems`
	// MoveToFolder
	RetentionActionMoveToFolder RetentionAction = `MoveToFolder`
	// None
	RetentionActionNone RetentionAction = `None`
	// PermanentlyDelete
	RetentionActionPermanentlyDelete RetentionAction = `PermanentlyDelete`
)
