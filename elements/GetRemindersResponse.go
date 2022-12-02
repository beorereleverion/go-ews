package elements

// The GetRemindersResponse element specifies the response to a GetReminders request.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/getremindersresponse
type GetRemindersResponse struct {
	// The Reminders element specifies the reminders returned in the response to a GetReminders request.
	Reminders *Reminders `xml:"m:Reminders"`
}
