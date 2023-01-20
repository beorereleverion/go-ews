package elements

// The Sender element represents the e-mail address for the sender of the message.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/sender-emailaddresstype
import "encoding/xml"

type SenderEmailAddressType struct {
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

func (S *SenderEmailAddressType) SetForMarshal() {
	S.XMLName.Local = "t:Sender"
}

func (S *SenderEmailAddressType) GetSchema() *Schema {
	return &SchemaTypes
}
