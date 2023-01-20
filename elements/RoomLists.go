package elements

// The RoomLists element is a list of one or more addresses that represent a list of meeting rooms.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/roomlists
import "encoding/xml"

type RoomLists struct {
	XMLName xml.Name

	// The Address element represents a fully resolved e-mail address.
	Address *AddressEmailAddressType `xml:"Address"`
}

func (R *RoomLists) SetForMarshal() {
	R.XMLName.Local = "m:RoomLists"
}

func (R *RoomLists) GetSchema() *Schema {
	return &SchemaMessages
}
