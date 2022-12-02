package elements

// The ProposedStart (AttendeeType) element specifies an attendee's proposed start time for a meeting.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/proposedstart-attendeetype
import "time"

type ProposedStartAttendeeType time.Time
