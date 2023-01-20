package elements

// The ReminderDueBy element represents the date and time when the event occurs. This is used by the ReminderMinutesBeforeStart element to determine when the reminder is displayed.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/reminderdueby
import "time"

import "encoding/xml"

type ReminderDueBy struct {
	XMLName xml.Name
	TEXT    time.Time `xml:",chardata"`
}

func (R *ReminderDueBy) SetForMarshal() {
	R.XMLName.Local = "t:ReminderDueBy"
}

func (R *ReminderDueBy) GetSchema() *Schema {
	return &SchemaTypes
}
