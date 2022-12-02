package elements

// The Time element describes the time when the time changes between standard time and daylight saving time.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/time-timechangetype
import "time"

type TimeTimeChangeType time.Time
