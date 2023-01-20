package elements

// The DailyRecurrence element describes the frequency, in days, in which a calendar item or a task recurs.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/dailyrecurrence
import "encoding/xml"

type DailyRecurrence struct {
	XMLName xml.Name

	// The Interval element defines the interval between two consecutive recurring items.
	Interval *Interval `xml:"Interval"`
}

func (D *DailyRecurrence) SetForMarshal() {
	D.XMLName.Local = "t:DailyRecurrence"
}

func (D *DailyRecurrence) GetSchema() *Schema {
	return &SchemaTypes
}
