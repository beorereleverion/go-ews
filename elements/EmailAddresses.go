package elements

// The EmailAddresses element represents a collection of e-mail addresses for a contact.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/emailaddresses
type EmailAddresses struct {
	// The Entry element represents a single e-mail address for a contact.
	Entry *EntryEmailAddress `xml:"t:Entry"`
}
