package elements

// The IndividualAttendeeConflictData element contains a user's or contact's free/busy status for a time window that occurs at the same time as the suggested meeting time identified in the Suggestion element.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/individualattendeeconflictdata
type IndividualAttendeeConflictData struct {
	// The BusyType element represents the free/busy status set for a calendar event.
	BusyType *BusyType `xml:"t:BusyType"`
}
