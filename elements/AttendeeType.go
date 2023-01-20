package elements

// The AttendeeType element represents the type of attendee that is identified in the Email (EmailAddressType) element. This element is used in requests for meeting suggestions.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/attendeetype
import "encoding/xml"

type AttendeeType struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

const (
	// A mailbox user who is an optional attendee to the meeting.
	AttendeeTypeOptional string = `Optional`
	// The mailbox user and attendee who created the calendar item.
	AttendeeTypeOrganizer string = `Organizer`
	// A mailbox user who is a required attendee to the meeting.
	AttendeeTypeRequired string = `Required`
	// A resource such as a TV or projector that is scheduled for use in the meeting.
	AttendeeTypeResource string = `Resource`
	// A mailbox entity that represents a room resource used for the meeting.
	AttendeeTypeRoom string = `Room`
)

func (A *AttendeeType) SetForMarshal() {
	A.XMLName.Local = "t:AttendeeType"
}

func (A *AttendeeType) GetSchema() *Schema {
	return &SchemaTypes
}
