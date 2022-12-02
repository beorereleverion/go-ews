package elements

// The ProposedEnd (AttendeeType) element specifies an attendee's proposed end time for a meeting.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/proposedend-attendeetype
import "time"

type ProposedEndAttendeeType time.Time
