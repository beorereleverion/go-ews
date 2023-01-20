package elements

// The NewReminderTime element specifies a new time for a reminder.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/newremindertime
import "encoding/xml"

type NewReminderTime struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (N *NewReminderTime) SetForMarshal() {
	N.XMLName.Local = "t:NewReminderTime"
}

func (N *NewReminderTime) GetSchema() *Schema {
	return &SchemaTypes
}
