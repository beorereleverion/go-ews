package elements

// The NewBodyContent element represents the new body content of a message.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/newbodycontent
type NewBodyContent struct {
	// Represents the actual body content of a message.
	BodyType *string `xml:"BodyType,attr"`
}
