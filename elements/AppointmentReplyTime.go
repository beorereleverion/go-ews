package elements

// The AppointmentReplyTime element represents the date and time that an attendee replied to a meeting request.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/appointmentreplytime
import "time"

type AppointmentReplyTime time.Time
