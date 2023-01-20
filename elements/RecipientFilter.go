package elements

// The RecipientFilter element represents a recipient address to use with the specified message tracking report.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/recipientfilter
import "encoding/xml"

type RecipientFilter struct {
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

func (R *RecipientFilter) SetForMarshal() {
	R.XMLName.Local = "m:RecipientFilter"
}

func (R *RecipientFilter) GetSchema() *Schema {
	return &SchemaMessages
}
