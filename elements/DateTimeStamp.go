package elements

// The DateTimeStamp element indicates the date and time that an instance of a calendar object was created.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/datetimestamp
import "time"

type DateTimeStamp time.Time
