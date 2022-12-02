package elements

// The GlobalItemClasses element contains a list of item classes that represents all the item classes of the conversation items in a mailbox.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/globalitemclasses
type GlobalItemClasses struct {
	// The ItemClass element represents the message class of an item.
	ItemClass *ItemClass `xml:"t:ItemClass"`
}
