package elements

// The AssociatedCalendarItemId element represents the calendar item that is associated with a MeetingMessage, MeetingRequest, MeetingResponse, MeetingCancellation, or ReminderMessageData.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/associatedcalendaritemid
type AssociatedCalendarItemId struct {
	// Identifies a specific version of the calendar item that is associated with a meeting.
	ChangeKey *string `xml:"ChangeKey,attr"`
	// Identifies the calendar item that is associated with meeting.
	Id *string `xml:"Id,attr"`
}
