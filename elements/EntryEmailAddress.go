package elements

// The Entry element represents a single e-mail address for a contact.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/entry-emailaddress
import "encoding/xml"

type EntryEmailAddress struct {
	XMLName xml.Name

	// Identifies the e-mail address.The following are the possible values for this attribute:- EmailAddress1  - EmailAddress2  - EmailAddress3   This attribute is required.
	Key *string `xml:"Key,attr"`
	// Defines the mailbox type of a mailbox user. This attribute is optional.
	MailboxType *string `xml:"MailboxType,attr"`
	// Defines the name of the mailbox user. This attribute is optional.
	Name *string `xml:"Name,attr"`
	// Defines the routing that is used for the mailbox. The default is SMTP. This attribute is optional.
	RoutingType *string `xml:"RoutingType,attr"`
}

func (E *EntryEmailAddress) SetForMarshal() {
	E.XMLName.Local = "t:Entry"
}

func (E *EntryEmailAddress) GetSchema() *Schema {
	return &SchemaTypes
}
