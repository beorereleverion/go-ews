package elements

// The GetRooms element is the root element in a request to get a list of rooms within a particular room list.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/getrooms
type GetRooms struct {
	// The RoomList element represents an e-mail address that identifies a list of meeting rooms.
	RoomList *RoomList `xml:"m:RoomList"`
}
