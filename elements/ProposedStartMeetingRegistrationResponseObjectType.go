package elements

// The ProposedStart (MeetingRegistrationResponseObjectType) element specifies an attendee's proposed new start time for a meeting.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/proposedstart-meetingregistrationresponseobjecttype
import "time"

type ProposedStartMeetingRegistrationResponseObjectType time.Time
