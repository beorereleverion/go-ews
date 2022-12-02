package elements

// The ItemClasses element contains a list of item classes that represents all the item classes of the conversation items in the current folder.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/itemclasses-arrayofitemclasstype
type ItemClassesArrayOfItemClassType struct {
	// The ItemClass element represents the message class of an item.
	ItemClass *ItemClass `xml:"t:ItemClass"`
}
