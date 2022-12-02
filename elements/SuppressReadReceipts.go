package elements

// The SuppressReadReceipts element indicates whether read receipts should be suppressed.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/suppressreadreceipts
type SuppressReadReceipts bool

const (
	// false
	SuppressReadReceiptsfalse SuppressReadReceipts = false
	// true
	SuppressReadReceiptstrue SuppressReadReceipts = true
)
