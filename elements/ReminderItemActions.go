package elements

// The ReminderItemActions element specifies the actions for reminder items.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/reminderitemactions
type ReminderItemActions struct {
	// The ReminderItemAction element specifies the action for a reminder item.
	ReminderItemAction *ReminderItemAction `xml:"t:ReminderItemAction"`
}
