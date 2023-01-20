package elements

// The ReminderGroup element specifies whether the reminder is for a calendar item or a task.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/remindergroup
import "encoding/xml"

type ReminderGroup struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

const (
	// Calendar
	ReminderGroupCalendar string = `Calendar`
	// Task
	ReminderGroupTask string = `Task`
)

func (R *ReminderGroup) SetForMarshal() {
	R.XMLName.Local = "t:ReminderGroup"
}

func (R *ReminderGroup) GetSchema() *Schema {
	return &SchemaTypes
}
