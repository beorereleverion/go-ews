package elements

// The PurportedSender element contains contact information for the alleged sender of an e-mail message.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/purportedsender
import "encoding/xml"

type PurportedSender struct {
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

func (P *PurportedSender) SetForMarshal() {
	P.XMLName.Local = "m:PurportedSender"
}

func (P *PurportedSender) GetSchema() *Schema {
	return &SchemaMessages
}
