package elements

// The RoomList element represents an e-mail address that identifies a list of meeting rooms.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/roomlist
import "encoding/xml"

type RoomList struct {
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

func (R *RoomList) SetForMarshal() {
	R.XMLName.Local = "m:RoomList"
}

func (R *RoomList) GetSchema() *Schema {
	return &SchemaMessages
}
