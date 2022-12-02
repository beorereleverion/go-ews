package elements

// The Room element represents a meeting room.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/room
type Room struct {
	// The Id element identifies a meeting room within the Exchange server organization.
	Id *IdEmailAddressType `xml:"t:Id"`
}
