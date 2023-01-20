package elements

// The ReminderNextTime element specifies the date and time for the next reminder.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/remindernexttime
import "time"

import "encoding/xml"

type ReminderNextTime struct {
	XMLName xml.Name
	TEXT    time.Time `xml:",chardata"`
}

func (R *ReminderNextTime) SetForMarshal() {
	R.XMLName.Local = "t:ReminderNextTime"
}

func (R *ReminderNextTime) GetSchema() *Schema {
	return &SchemaTypes
}
