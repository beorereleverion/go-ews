package elements

// The ExpandDL element defines a request to expand a distribution list.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/expanddl
type ExpandDL struct {
	// The Mailbox element identifies a mail-enabled Active Directory object.
	Mailbox *Mailbox `xml:"t:Mailbox"`
}
