package elements

// The EndDateTime element specifies the end date and time for a rule or a search.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/enddatetime
import "time"

type EndDateTime time.Time
