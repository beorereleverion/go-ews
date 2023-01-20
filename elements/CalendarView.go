package elements

// The CalendarView element defines a FindItem operation as returning calendar items in a set as they appear in a calendar.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/calendarview
import "encoding/xml"

type CalendarView struct {
	XMLName xml.Name

	// Identifies the end of a time span queried for calendar items. All calendar items that have a start time that is on or after EndDate will not be returned. The value of EndDate can be specified in UTC format, as in 2006-02-02T12:00:00Z, or in a format where local time and time zone offset is specified, as in 2006-02-02T04:00:00-08:00.  EndDate must be greater than or equal to StartDate; otherwise an error is returned. This attribute is required.
	EndDate *string `xml:"EndDate,attr"`
	// Describes the maximum number of results to return in the FindItem response.
	MaxEntriesReturned *string `xml:"MaxEntriesReturned,attr"`
	// Identifies the start of a time span queried for calendar items. All calendar items that have an end time that is before StartDate will not be returned. The value of StartDate can be specified in coordinated universal time (UTC) format, as in 2006-01-02T12:00:00Z, or in a format where local time and time zone offset is specified, as in 2006-01-02T04:00:00-08:00.  This attribute is required.
	StartDate *string `xml:"StartDate,attr"`
}

func (C *CalendarView) SetForMarshal() {
	C.XMLName.Local = "m:CalendarView"
}

func (C *CalendarView) GetSchema() *Schema {
	return &SchemaMessages
}
