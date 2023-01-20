package elements

// The Email element represents the mailbox user for a GetUserAvailability query.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/email-emailaddresstype
import "encoding/xml"

type EmailEmailAddressType struct {
	XMLName xml.Name

	// The Address element represents the e-mail address of the mailbox user.
	Address *Addressstring `xml:"Address"`
	// The Name element represents the display name of the mailbox user.
	Name *NameEmailAddress `xml:"Name"`
	// The RoutingType element represents the routing protocol for the recipient.
	RoutingType *RoutingTypeEmailAddress `xml:"RoutingType"`
}

func (E *EmailEmailAddressType) SetForMarshal() {
	E.XMLName.Local = "t:Email"
}

func (E *EmailEmailAddressType) GetSchema() *Schema {
	return &SchemaTypes
}
