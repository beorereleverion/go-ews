package elements

// The ReminderMessageData element specifies the data in a reminder message.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/remindermessagedata
type ReminderMessageData struct {
	// The AssociatedCalendarItemId element represents the calendar item that is associated with a MeetingMessage, MeetingRequest, MeetingResponse, MeetingCancellation, or ReminderMessageData.
	AssociatedCalendarItemId *AssociatedCalendarItemId `xml:"t:AssociatedCalendarItemId"`
	// The EndTime element represents the end of the time span to query for reminders.
	EndTime *EndTimeReminderMessageDataType `xml:"t:EndTime"`
	// The Location element represents the location of a meeting, appointment, or persona.
	Location *Location `xml:"t:Location"`
	// The ReminderText element specifies the text of a reminder message.
	ReminderText *ReminderText `xml:"t:ReminderText"`
	// The StartTime (ReminderMessageDataType) element specifies the starting time of the item that the reminder is for.
	StartTime *StartTimeReminderMessageDataType `xml:"t:StartTime"`
}
