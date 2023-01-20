package elements

// The EmailAddress element specifies the fully resolved SMTP address for the site mailbox or the associated persona.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/emailaddress-emailaddresstype
import "encoding/xml"

type EmailAddressEmailAddressType struct {
	XMLName xml.Name

	// The EmailAddress element defines the primary SMTP address of a mailbox user.
	EmailAddress *EmailAddressNonEmptyStringType `xml:"EmailAddress"`
	// The ItemId element contains the unique identifier and change key of an item in the Exchange store.
	ItemId *ItemId `xml:"ItemId"`
	// The MailboxType element represents the type of mailbox that is represented by the e-mail address.
	MailboxType *MailboxType `xml:"MailboxType"`
	// The Name element specifies a search refiner name.
	Name *Namestring `xml:"Name"`
	// The RoutingType element defines the address type for the mailbox.
	RoutingType *RoutingTypeEmailAddressType `xml:"RoutingType"`
}

func (E *EmailAddressEmailAddressType) SetForMarshal() {
	E.XMLName.Local = "t:EmailAddress"
}

func (E *EmailAddressEmailAddressType) GetSchema() *Schema {
	return &SchemaTypes
}
