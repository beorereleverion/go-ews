package elements

// The ActionType element specifies the action to take on the reminder.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/actiontype-reminderactiontype
type ActionTypeReminderActionType string

const (
	// Dismiss
	ActionTypeReminderActionTypeDismiss ActionTypeReminderActionType = `Dismiss`
	// Snooze
	ActionTypeReminderActionTypeSnooze ActionTypeReminderActionType = `Snooze`
)
