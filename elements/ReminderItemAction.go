package elements

// The ReminderItemAction element specifies the action for a reminder item.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/reminderitemaction
type ReminderItemAction struct {
	// The ActionType element specifies the action to take on the reminder.
	ActionType *ActionTypeReminderActionType `xml:"t:ActionType"`
	// The ItemId element contains the unique identifier and change key of an item in the Exchange store.
	ItemId *ItemId `xml:"t:ItemId"`
	// The NewReminderTime element specifies a new time for a reminder.
	NewReminderTime *NewReminderTime `xml:"t:NewReminderTime"`
}
