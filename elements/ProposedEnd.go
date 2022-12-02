package elements

// The ProposedEnd element specifies the proposed end time of a meeting.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/proposedend
import "time"

type ProposedEnd time.Time
