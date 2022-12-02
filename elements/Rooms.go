package elements

// The Rooms element is a list of one or more elements that represent meeting rooms.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/rooms
type Rooms struct {
	// The Room element represents a meeting room.
	Room *Room `xml:"t:Room"`
}
