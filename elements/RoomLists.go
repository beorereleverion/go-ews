package elements

// The RoomLists element is a list of one or more addresses that represent a list of meeting rooms.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/roomlists
type RoomLists struct {
	// The Address element represents a fully resolved e-mail address.
	Address *AddressEmailAddressType `xml:"t:Address"`
}
