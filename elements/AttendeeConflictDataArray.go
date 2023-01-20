package elements

// The AttendeeConflictDataArray element contains an array of conflict data for queried attendees identified in the GetUserAvailability operation.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/attendeeconflictdataarray
import "encoding/xml"

type AttendeeConflictDataArray struct {
	XMLName xml.Name

	// The GroupAttendeeConflictData element contains aggregate conflict information about the number of users who are available, the number of users who have conflicts, and the number of users who do not have availability information in a distribution list for a suggested meeting time.
	GroupAttendeeConflictData *GroupAttendeeConflictData `xml:"GroupAttendeeConflictData"`
	// The IndividualAttendeeConflictData element contains a user's or contact's free/busy status for a time window that occurs at the same time as the suggested meeting time identified in the Suggestion element.
	IndividualAttendeeConflictData *IndividualAttendeeConflictData `xml:"IndividualAttendeeConflictData"`
	// The Suggestion element represents a single meeting suggestion.
	Suggestion *Suggestion `xml:"Suggestion"`
	// The TooBigGroupAttendeeConflictData element represents an attendee that was resolved as a distribution list but the distribution list was too large to expand.
	TooBigGroupAttendeeConflictData *TooBigGroupAttendeeConflictData `xml:"TooBigGroupAttendeeConflictData"`
	// The UnknownAttendeeConflictData element represents an unresolvable attendee or an attendee that is not a user, distribution list, or contact.
	UnknownAttendeeConflictData *UnknownAttendeeConflictData `xml:"UnknownAttendeeConflictData"`
}

func (A *AttendeeConflictDataArray) SetForMarshal() {
	A.XMLName.Local = "t:AttendeeConflictDataArray"
}

func (A *AttendeeConflictDataArray) GetSchema() *Schema {
	return &SchemaTypes
}
