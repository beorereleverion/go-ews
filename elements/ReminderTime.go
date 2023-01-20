package elements

// The ReminderTime element specifies the time for the reminder to occur.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/remindertime
import "time"

import "encoding/xml"

type ReminderTime struct {
	XMLName xml.Name
	TEXT    time.Time `xml:",chardata"`
}

func (R *ReminderTime) SetForMarshal() {
	R.XMLName.Local = "t:ReminderTime"
}

func (R *ReminderTime) GetSchema() *Schema {
	return &SchemaTypes
}
