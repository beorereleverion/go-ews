package elements

// The Birthday element represents the birth date of a contact.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/birthday
import "time"

type Birthday time.Time
