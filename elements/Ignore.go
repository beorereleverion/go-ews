package elements

// The Ignore element identifies items to skip during synchronization.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/ignore
type Ignore struct {
	// The ItemId element contains the unique identifier and change key of an item in the Exchange store.
	ItemId *ItemId `xml:"t:ItemId"`
}
