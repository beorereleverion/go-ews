package elements

// The Id element identifies a meeting room within the Exchange server organization.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/id-emailaddresstype
import "encoding/xml"

type IdEmailAddressType struct {
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

func (I *IdEmailAddressType) SetForMarshal() {
	I.XMLName.Local = "t:Id"
}

func (I *IdEmailAddressType) GetSchema() *Schema {
	return &SchemaTypes
}
