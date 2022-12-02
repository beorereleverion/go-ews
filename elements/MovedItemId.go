package elements

// The MovedItemId element specifies the identifier of the item that was moved by the MarkAsJunk operation.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/moveditemid
type MovedItemId struct {
	// The value of the ChangeKey attribute is the change key of the moved item. The change key changes after the item is moved by the MarkAsJunk operation.
	ChangeKey *string `xml:"ChangeKey,attr"`
	// The value of the Id attribute is the item identifier of the item that is moved by the MarkAsJunk operation. The item identifier will remain the same after the move.
	Id *string `xml:"Id,attr"`
}
