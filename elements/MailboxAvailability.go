package elements

// The Mailbox element represents the mailbox user for a SetUserOofSettings or GetUserOofSettings request.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/mailbox-availability
import "encoding/xml"

type MailboxAvailability struct {
	XMLName xml.Name

	// The Address element represents the e-mail address of the mailbox user.
	Address *Addressstring `xml:"Address"`
	// The Name element represents the display name of the mailbox user.
	Name *NameEmailAddress `xml:"Name"`
	// The RoutingType element represents the routing protocol for the recipient.
	RoutingType *RoutingTypeEmailAddress `xml:"RoutingType"`
}

func (M *MailboxAvailability) SetForMarshal() {
	M.XMLName.Local = "t:Mailbox"
}

func (M *MailboxAvailability) GetSchema() *Schema {
	return &SchemaTypes
}
