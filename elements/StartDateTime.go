package elements

// The StartDateTime element specifies the start date and time for a rule or a search.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/startdatetime
import "time"

type StartDateTime time.Time
