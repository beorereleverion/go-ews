package elements

// The IsReadReceipt element indicates whether incoming messages must be read receipts in order for the condition or exception to apply.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/isreadreceipt
type IsReadReceipt bool

const (
	// false
	IsReadReceiptfalse IsReadReceipt = false
	// true
	IsReadReceipttrue IsReadReceipt = true
)
