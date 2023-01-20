package elements

// The Reminders element specifies the reminders returned in the response to a GetReminders request.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/reminders
import "encoding/xml"

type Reminders struct {
	XMLName xml.Name

	// The Reminder element specifies a reminder for a task or a calendar item.
	Reminder *Reminder `xml:"Reminder"`
}

func (R *Reminders) SetForMarshal() {
	R.XMLName.Local = "m:Reminders"
}

func (R *Reminders) GetSchema() *Schema {
	return &SchemaMessages
}
