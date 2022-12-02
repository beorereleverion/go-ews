package elements

// The DeliveryRestricted element indicates whether delivery restrictions will prevent the sender's message from reaching the recipient.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/deliveryrestricted
type DeliveryRestricted bool

const (
	// false
	DeliveryRestrictedfalse DeliveryRestricted = false
	// true
	DeliveryRestrictedtrue DeliveryRestricted = true
)
