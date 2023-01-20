package elements

// The DayOfWeek element contains the list of working days scheduled for the mailbox user.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/dayofweek-workingperiod
import "encoding/xml"

type DayOfWeekWorkingPeriod struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (D *DayOfWeekWorkingPeriod) SetForMarshal() {
	D.XMLName.Local = "t:DayOfWeek"
}

func (D *DayOfWeekWorkingPeriod) GetSchema() *Schema {
	return &SchemaTypes
}
