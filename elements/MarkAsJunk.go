package elements

// The MarkAsJunk element specifies the request to move an item to the junk mail folder and to add the sender to the blocked sender list.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/markasjunk
import "encoding/xml"

type MarkAsJunk struct {
	XMLName xml.Name

	// The ItemIds element contains the unique identities of items, occurrence items, and recurring master items that are used to delete, send, get, move, or copy items in the Exchange store.
	ItemIds *ItemIds `xml:"ItemIds"`
	// A text value of true for the IsJunk attribute indicates that the email sender is added to the blocked sender list. A value of false indicates that the email sender is removed from the blocked sender list, if the email sender is already on the list.
	IsJunk *string `xml:"IsJunk,attr"`
	// A text value of true for the MoveItem attribute indicates that the item is moved to the default junk mail folder. A value of false indicates that the item is not moved to the default junk mail folder.
	MoveItem *string `xml:"MoveItem,attr"`
}

func (M *MarkAsJunk) SetForMarshal() {
	M.XMLName.Local = "m:MarkAsJunk"
}

func (M *MarkAsJunk) GetSchema() *Schema {
	return &SchemaMessages
}
