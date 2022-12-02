package elements

// The InternalReply element contains the out of office (OOF) response sent to other users in the user's domain or trusted domains.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/internalreply
type InternalReply struct {
	// The Message element contains the out of Office (OOF) response.
	Message *MessageAvailability `xml:"t:Message"`
}
