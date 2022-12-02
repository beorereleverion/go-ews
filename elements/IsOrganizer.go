package elements

// The IsOrganizer element specifies a Boolean value that indicates whether this person is the organizer of the meeting.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/isorganizer
type IsOrganizer bool

const (
	// false
	IsOrganizerfalse IsOrganizer = false
	// true
	IsOrganizertrue IsOrganizer = true
)
