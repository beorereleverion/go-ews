package elements

// The ProposedEnd (MeetingRegistrationResponseObjectType) element specifies an attendee's proposed new end time for a meeting.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/proposedend-meetingregistrationresponseobjecttype
import "time"

type ProposedEndMeetingRegistrationResponseObjectType time.Time
