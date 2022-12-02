package elements

// The ImAddresses element represents a collection of instant messaging addresses for a contact.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/imaddresses
type ImAddresses struct {
	// The Entry element represents an instant messaging (IM) address for a contact.
	Entry *EntryIMAddress `xml:"t:Entry"`
}
