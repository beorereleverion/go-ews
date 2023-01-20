package elements

// The ReminderText element specifies the text of a reminder message.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/remindertext
import "encoding/xml"

type ReminderText struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (R *ReminderText) SetForMarshal() {
	R.XMLName.Local = "t:ReminderText"
}

func (R *ReminderText) GetSchema() *Schema {
	return &SchemaTypes
}
