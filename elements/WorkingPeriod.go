package elements

// The WorkingPeriod element contains the work week days and hours of the mailbox user.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/workingperiod
import "encoding/xml"

type WorkingPeriod struct {
	XMLName xml.Name

	// The DayOfWeek element contains the list of working days scheduled for the mailbox user.
	DayOfWeek *DayOfWeekWorkingPeriod `xml:"DayOfWeek"`
	// The EndTimeInMinutes element represents the end of the working day for a mailbox user.
	EndTimeInMinutes *EndTimeInMinutes `xml:"EndTimeInMinutes"`
	// The StartTimeInMinutes element represents the start of the working day for a mailbox user.
	StartTimeInMinutes *StartTimeInMinutes `xml:"StartTimeInMinutes"`
}

func (W *WorkingPeriod) SetForMarshal() {
	W.XMLName.Local = "t:WorkingPeriod"
}

func (W *WorkingPeriod) GetSchema() *Schema {
	return &SchemaTypes
}
