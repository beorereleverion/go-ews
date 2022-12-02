package elements

// The LastModifiedTime element indicates when an item was last modified. This element is read-only.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/lastmodifiedtime
import "time"

type LastModifiedTime time.Time
