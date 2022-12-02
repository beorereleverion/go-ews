package elements

// The Reminders element specifies the reminders returned in the response to a GetReminders request.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/reminders
type Reminders struct {
	// The Reminder element specifies a reminder for a task or a calendar item.
	Reminder *Reminder `xml:"t:Reminder"`
}
