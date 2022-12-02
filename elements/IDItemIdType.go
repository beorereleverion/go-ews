package elements

// The Id element specifies the identifier of an item.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/id-itemidtype
type IDItemIdType struct {
	// The text value of the ChangeKey attribute is the change key of the item.
	ChangeKey *string `xml:"ChangeKey,attr"`
	// The text value of the Id attribute is the identifier of the item.
	Id *string `xml:"Id,attr"`
}
