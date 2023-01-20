package elements

// The AdjacentMeetings element identifies all calendar items that are adjacent to a meeting time.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/adjacentmeetings
import "encoding/xml"

type AdjacentMeetings struct {
	XMLName xml.Name

	// The CalendarItem element represents an Exchange calendar item.
	CalendarItem *CalendarItem `xml:"CalendarItem"`
}

func (A *AdjacentMeetings) SetForMarshal() {
	A.XMLName.Local = "t:AdjacentMeetings"
}

func (A *AdjacentMeetings) GetSchema() *Schema {
	return &SchemaTypes
}
