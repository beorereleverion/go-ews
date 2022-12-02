package elements

// The End element specifies the changes made to a meeting end time when a meeting update occurs.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/end-changehighlightstype
import "time"

type EndChangeHighlightsType time.Time
