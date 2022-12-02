package elements

// The UniqueBody element represents an HTML fragment or plain text which represents the unique body of this conversation.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/uniquebody
type UniqueBody struct {
	// Describes how the item body is stored in the item.
	BodyType *string `xml:"BodyType,attr"`
}
