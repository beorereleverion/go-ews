package elements

// The FederatedDeliveryMailbox element represents the mailbox to which a cross-premise message was sent.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/federateddeliverymailbox
import "encoding/xml"

type FederatedDeliveryMailbox struct {
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

func (F *FederatedDeliveryMailbox) SetForMarshal() {
	F.XMLName.Local = "m:FederatedDeliveryMailbox"
}

func (F *FederatedDeliveryMailbox) GetSchema() *Schema {
	return &SchemaMessages
}
