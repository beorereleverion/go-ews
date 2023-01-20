package elements

// The AssociatedCalendarItemId element represents the calendar item that is associated with a MeetingMessage, MeetingRequest, MeetingResponse, MeetingCancellation, or ReminderMessageData.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/associatedcalendaritemid
import "encoding/xml"

type AssociatedCalendarItemId struct {
	XMLName xml.Name

	// Identifies a specific version of the calendar item that is associated with a meeting.
	ChangeKey *string `xml:"ChangeKey,attr"`
	// Identifies the calendar item that is associated with meeting.
	Id *string `xml:"Id,attr"`
}

func (A *AssociatedCalendarItemId) SetForMarshal() {
	A.XMLName.Local = "t:AssociatedCalendarItemId"
}

func (A *AssociatedCalendarItemId) GetSchema() *Schema {
	return &SchemaTypes
}
