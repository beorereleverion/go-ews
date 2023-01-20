package elements

// The UID (ReminderType) element identifies the calendar item associated with a reminder.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/uid-remindertype
import "encoding/xml"

type UIDReminderType struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (U *UIDReminderType) SetForMarshal() {
	U.XMLName.Local = "t:UID"
}

func (U *UIDReminderType) GetSchema() *Schema {
	return &SchemaTypes
}
