package elements

// The ItemChanges element contains an array of ItemChange elements that identify items and the updates to apply to the items.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/itemchanges
type ItemChanges struct {
	// The ItemChange element contains an item identifier and the updates to apply to the item.
	ItemChange *ItemChange `xml:"t:ItemChange"`
}
