package elements

// The ExternalReply element contains the out of office (OOF) response that is sent to addresses outside the recipient's domain or trusted domains.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/externalreply
type ExternalReply struct {
	// The Message element contains the out of Office (OOF) response.
	Message *MessageAvailability `xml:"t:Message"`
}
