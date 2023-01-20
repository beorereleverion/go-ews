package elements

// The ConflictingMeetings element identifies all calendar items that conflict with a meeting time.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/conflictingmeetings
import "encoding/xml"

type ConflictingMeetings struct {
	XMLName xml.Name

	// The CalendarItem element represents an Exchange calendar item.
	CalendarItem *CalendarItem `xml:"CalendarItem"`
}

func (C *ConflictingMeetings) SetForMarshal() {
	C.XMLName.Local = "t:ConflictingMeetings"
}

func (C *ConflictingMeetings) GetSchema() *Schema {
	return &SchemaTypes
}
