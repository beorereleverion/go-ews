package elements

// The DeliveryStatus element specifies the status for a message.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/deliverystatus
type DeliveryStatus string

const (
	// Delivered
	DeliveryStatusDelivered DeliveryStatus = `Delivered`
	// Pending
	DeliveryStatusPending DeliveryStatus = `Pending`
	// Read
	DeliveryStatusRead DeliveryStatus = `Read`
	// Transferred
	DeliveryStatusTransferred DeliveryStatus = `Transferred`
	// Unsuccessful
	DeliveryStatusUnsuccessful DeliveryStatus = `Unsuccessful`
)
