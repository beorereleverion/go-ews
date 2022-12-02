package elements

// The PerformReminderAction element specifies a request to perform a reminder action.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/performreminderaction
type PerformReminderAction struct {
	// The ReminderItemActions element specifies the actions for reminder items.
	ReminderItemActions *ReminderItemActions `xml:"m:ReminderItemActions"`
}
