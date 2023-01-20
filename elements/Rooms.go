package elements

// The Rooms element is a list of one or more elements that represent meeting rooms.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/rooms
import "encoding/xml"

type Rooms struct {
	XMLName xml.Name

	// The Room element represents a meeting room.
	Room *Room `xml:"Room"`
}

func (R *Rooms) SetForMarshal() {
	R.XMLName.Local = "m:Rooms"
}

func (R *Rooms) GetSchema() *Schema {
	return &SchemaMessages
}
