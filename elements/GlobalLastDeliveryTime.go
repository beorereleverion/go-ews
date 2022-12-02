package elements

// The GlobalLastDeliveryTime element contains the delivery time of the message that was last received in this conversation across all folders in the mailbox.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/globallastdeliverytime
import "time"

type GlobalLastDeliveryTime time.Time
