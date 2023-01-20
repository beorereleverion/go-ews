package elements

// The Room element represents a meeting room.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/room
import "encoding/xml"

type Room struct {
	XMLName xml.Name

	// The Id element identifies a meeting room within the Exchange server organization.
	Id *IdEmailAddressType `xml:"Id"`
}

func (R *Room) SetForMarshal() {
	R.XMLName.Local = "t:Room"
}

func (R *Room) GetSchema() *Schema {
	return &SchemaTypes
}
