package elements

// The OriginalStart element represents the original start time of a calendar item.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/originalstart
import "time"

type OriginalStart time.Time
