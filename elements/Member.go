package elements

// The Member element represents a member of a distribution list.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/member-ex15websvcsotherref
type Member struct {
	// The Mailbox element identifies a mail-enabled Active Directory object.
	Mailbox *Mailbox `xml:"t:Mailbox"`
	// The Status element provides information about the status of a distribution list member on the server.
	Status *StatusMemberStatusType `xml:"t:Status"`
	// Provides a unique identifier for the distribution list member. This attribute is optional.
	Key *string `xml:"Key,attr"`
}
