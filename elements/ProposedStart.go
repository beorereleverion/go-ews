package elements

// The ProposedStart element specifies the proposed start time of a meeting.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/proposedstart
import "time"

type ProposedStart time.Time
