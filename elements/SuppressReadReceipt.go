package elements

// The SuppressReadReceipt element is used to suppress read receipts.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/suppressreadreceipt
type SuppressReadReceipt struct {
	// The ReferenceItemId element identifies the item to which the response object refers.
	ReferenceItemId *ReferenceItemId `xml:"t:ReferenceItemId"`
}
