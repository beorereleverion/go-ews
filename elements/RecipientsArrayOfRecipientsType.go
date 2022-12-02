package elements

// The Recipients element represents a collection of recipients that receive a copy of the message.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/recipients-arrayofrecipientstype
type RecipientsArrayOfRecipientsType struct {
	// The Mailbox element identifies a mail-enabled Active Directory object.
	Mailbox *Mailbox `xml:"t:Mailbox"`
}
