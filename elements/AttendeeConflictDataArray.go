package elements

// The AttendeeConflictDataArray element contains an array of conflict data for queried attendees identified in the GetUserAvailability operation.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/attendeeconflictdataarray
type AttendeeConflictDataArray struct {
	// The GroupAttendeeConflictData element contains aggregate conflict information about the number of users who are available, the number of users who have conflicts, and the number of users who do not have availability information in a distribution list for a suggested meeting time.
	GroupAttendeeConflictData *GroupAttendeeConflictData `xml:"t:GroupAttendeeConflictData"`
	// The IndividualAttendeeConflictData element contains a user's or contact's free/busy status for a time window that occurs at the same time as the suggested meeting time identified in the Suggestion element.
	IndividualAttendeeConflictData *IndividualAttendeeConflictData `xml:"t:IndividualAttendeeConflictData"`
	// The Suggestion element represents a single meeting suggestion.
	Suggestion *Suggestion `xml:"t:Suggestion"`
	// The TooBigGroupAttendeeConflictData element represents an attendee that was resolved as a distribution list but the distribution list was too large to expand.
	TooBigGroupAttendeeConflictData *TooBigGroupAttendeeConflictData `xml:"t:TooBigGroupAttendeeConflictData"`
	// The UnknownAttendeeConflictData element represents an unresolvable attendee or an attendee that is not a user, distribution list, or contact.
	UnknownAttendeeConflictData *UnknownAttendeeConflictData `xml:"t:UnknownAttendeeConflictData"`
}
