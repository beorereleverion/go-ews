package elements

// The ReplyBody element contains an Out of Office (OOF) message and the language used for the message.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/replybody
type ReplyBody struct {
	// The Message element contains the out of Office (OOF) response.
	Message *MessageAvailability `xml:"t:Message"`
}
