package elements

// The CreationTime element specifies when the persona was created.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/creationtime
import "time"

type CreationTime time.Time
