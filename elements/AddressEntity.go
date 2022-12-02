package elements

// The AddressEntity element specifies a single address entity.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/addressentity
type AddressEntity struct {
	// The Address element represents the e-mail address of the mailbox user.
	Address *Addressstring `xml:"t:Address"`
	// The Position element specifies the position of an entity extracted from a message.
	Position *Position `xml:"t:Position"`
}
