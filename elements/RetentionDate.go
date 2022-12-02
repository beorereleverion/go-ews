package elements

// The RetentionDate element specifies the last date that an item must be retained.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/retentiondate
import "time"

type RetentionDate time.Time
