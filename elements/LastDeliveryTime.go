package elements

// The LastDeliveryTime element contains the delivery time of the message that was last received in this conversation in the current folder.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/lastdeliverytime
import "time"

type LastDeliveryTime time.Time
