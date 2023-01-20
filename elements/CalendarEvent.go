package elements

// The CalendarEvent element represents a unique calendar item occurrence.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/calendarevent
import "encoding/xml"

type CalendarEvent struct {
	XMLName xml.Name

	// The BusyType element represents the free/busy status set for a calendar event.
	BusyType *BusyType `xml:"BusyType"`
	// The CalendarEventDetails element provides additional information about a calendar event.
	CalendarEventDetails *CalendarEventDetails `xml:"CalendarEventDetails"`
	// The EndTime element represents the end of a time span.
	EndTime *EndTime `xml:"EndTime"`
	// The StartTime element represents the start of a time span.
	StartTime *StartTime `xml:"StartTime"`
}

func (C *CalendarEvent) SetForMarshal() {
	C.XMLName.Local = "t:CalendarEvent"
}

func (C *CalendarEvent) GetSchema() *Schema {
	return &SchemaTypes
}
