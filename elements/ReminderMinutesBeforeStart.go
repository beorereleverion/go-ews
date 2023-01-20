package elements

// The ReminderMinutesBeforeStart element represents the number of minutes before an event occurs when a reminder is displayed.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/reminderminutesbeforestart
import "encoding/xml"

type ReminderMinutesBeforeStart struct {
	XMLName xml.Name
	TEXT    int64 `xml:",chardata"`
}

func (R *ReminderMinutesBeforeStart) SetForMarshal() {
	R.XMLName.Local = "t:ReminderMinutesBeforeStart"
}

func (R *ReminderMinutesBeforeStart) GetSchema() *Schema {
	return &SchemaTypes
}
