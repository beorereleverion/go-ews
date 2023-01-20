package elements

// The Mailbox element identifies a mail-enabled Active Directory object.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/mailbox
import "encoding/xml"

type Mailbox struct {
	XMLName xml.Name

	// The EmailAddress element defines the primary SMTP address of a mailbox user.
	EmailAddress *EmailAddressNonEmptyStringType `xml:"EmailAddress"`
	// The ItemId element contains the unique identifier and change key of an item in the Exchange store.
	ItemId *ItemId `xml:"ItemId"`
	// The MailboxType element represents the type of mailbox that is represented by the e-mail address.
	MailboxType *MailboxType `xml:"MailboxType"`
	// The Name element represents the name of a mailbox user.
	Name *NameEmailAddressType `xml:"Name"`
	// The RoutingType element represents the routing protocol for the recipient.
	RoutingType *RoutingTypeEmailAddress `xml:"RoutingType"`
}

func (M *Mailbox) SetForMarshal() {
	M.XMLName.Local = "t:Mailbox"
}

func (M *Mailbox) GetSchema() *Schema {
	return &SchemaTypes
}
