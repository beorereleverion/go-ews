package elements

// The CurrentMeetingTime element represents the start time of a meeting that you want to update with a meeting time proposed by a meeting attendee.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/currentmeetingtime
import "time"

type CurrentMeetingTime time.Time
