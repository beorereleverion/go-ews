package elements

// The GetRooms element is the root element in a request to get a list of rooms within a particular room list.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/getrooms
import "encoding/xml"

type GetRooms struct {
	XMLName xml.Name

	// The RoomList element represents an e-mail address that identifies a list of meeting rooms.
	RoomList *RoomList `xml:"RoomList"`
}

func (G *GetRooms) SetForMarshal() {
	G.XMLName.Local = "m:GetRooms"
}

func (G *GetRooms) GetSchema() *Schema {
	return &SchemaMessages
}
