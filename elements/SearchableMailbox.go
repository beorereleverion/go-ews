package elements

// The SearchableMailbox element specifies a mailbox returned from a GetSearchableMailboxes request.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/searchablemailbox
import "encoding/xml"

type SearchableMailbox struct {
	XMLName xml.Name

	// The DisplayName element defines the display name of a folder, contact, distribution list, delegate user, location, or rule.
	DisplayName *DisplayNamestring `xml:"DisplayName"`
	// The ExternalEmailAddress element contains the external email address of the mailbox.
	ExternalEmailAddress *ExternalEmailAddress `xml:"ExternalEmailAddress"`
	// The Guid element specifies the globally unique identifier of the mailbox.
	Guid *Guid `xml:"Guid"`
	// The IsExternalMailbox element indicates whether the mailbox is external to the organization.
	IsExternalMailbox *IsExternalMailbox `xml:"IsExternalMailbox"`
	// The IsMembershipGroup element specifies a Boolean value that indicates whether the entity is a distribution group or a mailbox.
	IsMembershipGroup *IsMembershipGroup `xml:"IsMembershipGroup"`
	// The PrimarySmtpAddress element specifies the primary Simple Mail Transfer Protocol (SMTP) address of the mailbox.
	PrimarySmtpAddress *PrimarySmtpAddressstring `xml:"PrimarySmtpAddress"`
	// The ReferenceId element specifies the reference identifier for the mailbox.
	ReferenceId *ReferenceId `xml:"ReferenceId"`
}

func (S *SearchableMailbox) SetForMarshal() {
	S.XMLName.Local = "t:SearchableMailbox"
}

func (S *SearchableMailbox) GetSchema() *Schema {
	return &SchemaTypes
}
