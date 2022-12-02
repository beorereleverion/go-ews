package elements

// The GetReminders element specifies a request to get reminders.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/getreminders
type GetReminders struct {
	// The BeginTime element specifies the beginning of the time span to query for reminders.
	BeginTime *BeginTime `xml:"m:BeginTime"`
	// The EndTime element represents the end of the time span to query for reminders.
	EndTime *EndTimeReminderMessageDataType `xml:"t:EndTime"`
	// The MaxItems element specifies the maximum number of items to return in the request.
	MaxItems *MaxItems `xml:"m:MaxItems"`
	// The ReminderType element specifies the type of reminders to return.
	ReminderType *ReminderType `xml:"m:ReminderType"`
}
