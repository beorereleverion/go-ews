package elements

// The AttendeeType element represents the type of attendee that is identified in the Email (EmailAddressType) element. This element is used in requests for meeting suggestions.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/attendeetype
type AttendeeType string

const (
	// A mailbox user who is an optional attendee to the meeting.
	AttendeeTypeOptional AttendeeType = `Optional`
	// The mailbox user and attendee who created the calendar item.
	AttendeeTypeOrganizer AttendeeType = `Organizer`
	// A mailbox user who is a required attendee to the meeting.
	AttendeeTypeRequired AttendeeType = `Required`
	// A resource such as a TV or projector that is scheduled for use in the meeting.
	AttendeeTypeResource AttendeeType = `Resource`
	// A mailbox entity that represents a room resource used for the meeting.
	AttendeeTypeRoom AttendeeType = `Room`
)
