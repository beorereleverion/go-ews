package elements

// The AbsoluteYearlyRecurrence element represents a yearly recurrence pattern.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/absoluteyearlyrecurrence
import "encoding/xml"

type AbsoluteYearlyRecurrence struct {
	XMLName xml.Name

	// The DayOfMonth element describes the day in a month that a recurring item occurs.
	DayOfMonth *DayOfMonth `xml:"DayOfMonth"`
	// The Month element describes the month when a yearly recurring item occurs.
	Month *MonthItemRecurrence `xml:"Month"`
}

func (A *AbsoluteYearlyRecurrence) SetForMarshal() {
	A.XMLName.Local = "t:AbsoluteYearlyRecurrence"
}

func (A *AbsoluteYearlyRecurrence) GetSchema() *Schema {
	return &SchemaTypes
}
