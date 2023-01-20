package elements

// The ReminderType element specifies the type of reminders to return.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/remindertype
import "encoding/xml"

type ReminderType struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

const (
	// All
	ReminderTypeAll string = `All`
	// Current
	ReminderTypeCurrent string = `Current`
	// Old
	ReminderTypeOld string = `Old`
)

func (R *ReminderType) SetForMarshal() {
	R.XMLName.Local = "m:ReminderType"
}

func (R *ReminderType) GetSchema() *Schema {
	return &SchemaMessages
}
