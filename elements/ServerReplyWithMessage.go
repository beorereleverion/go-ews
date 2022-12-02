package elements

// The ServerReplyWithMessage element indicates the ID of the template message that is to be sent as a reply to incoming messages.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/serverreplywithmessage
type ServerReplyWithMessage struct {
	// The ItemId element contains the unique identifier and change key of an item in the Exchange store.
	ItemId *ItemId `xml:"t:ItemId"`
}
