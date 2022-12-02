package elements

// The UpdateItemInRecoverableItems element specifies a request to update an item in recoverable items.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/updateiteminrecoverableitems
type UpdateItemInRecoverableItems struct {
	// The Attachments element contains the items or files that are attached to an item in the Exchange store.
	Attachments *Attachments `xml:"t:Attachments"`
	// The ItemId element contains the unique identifier and change key of an item in the Exchange store.
	ItemId *ItemId `xml:"t:ItemId"`
	// The MakeItemImmutable element specifies a Boolean value that indicates whether an item should be made read-only.
	MakeItemImmutable *MakeItemImmutable `xml:"m:MakeItemImmutable"`
	// The Updates element contains a set of elements that define append, set, and delete changes to item properties.
	Updates *UpdatesItem `xml:"Updates"`
}
