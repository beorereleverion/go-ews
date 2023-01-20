package elements

// The AbsoluteMonthlyRecurrence element represents a monthly recurrence pattern.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/absolutemonthlyrecurrence
import "encoding/xml"

type AbsoluteMonthlyRecurrence struct {
	XMLName xml.Name

	// The DayOfMonth element describes the day in a month that a recurring item occurs.
	DayOfMonth *DayOfMonth `xml:"DayOfMonth"`
	// The Interval element defines the interval between two consecutive recurring items.
	Interval *Interval `xml:"Interval"`
}

func (A *AbsoluteMonthlyRecurrence) SetForMarshal() {
	A.XMLName.Local = "t:AbsoluteMonthlyRecurrence"
}

func (A *AbsoluteMonthlyRecurrence) GetSchema() *Schema {
	return &SchemaTypes
}
