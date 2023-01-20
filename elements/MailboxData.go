package elements

// The MailboxData element represents an individual mailbox user and options for the type of data to be returned about the mailbox user.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/mailboxdata
import "encoding/xml"

type MailboxData struct {
	XMLName xml.Name

	// The AttendeeType element represents the type of attendee that is identified in the Email (EmailAddressType) element. This element is used in requests for meeting suggestions.
	AttendeeType *AttendeeType `xml:"AttendeeType"`
	// The Email element represents the mailbox user for a GetUserAvailability query.
	Email *EmailEmailAddressType `xml:"Email"`
	// The ExcludeConflicts element specifies whether to return suggested times for calendar times that conflict among the attendees.
	ExcludeConflicts *ExcludeConflicts `xml:"ExcludeConflicts"`
}

func (M *MailboxData) SetForMarshal() {
	M.XMLName.Local = "t:MailboxData"
}

func (M *MailboxData) GetSchema() *Schema {
	return &SchemaTypes
}
