package elements

// The CalendarEventArray element contains a set of unique calendar item occurrences that represent the requested user's availability.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/calendareventarray
import "encoding/xml"

type CalendarEventArray struct {
	XMLName xml.Name

	// The CalendarEvent element represents a unique calendar item occurrence.
	CalendarEvent *CalendarEvent `xml:"CalendarEvent"`
}

func (C *CalendarEventArray) SetForMarshal() {
	C.XMLName.Local = "t:CalendarEventArray"
}

func (C *CalendarEventArray) GetSchema() *Schema {
	return &SchemaTypes
}
