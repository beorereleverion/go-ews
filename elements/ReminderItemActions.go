package elements

// The ReminderItemActions element specifies the actions for reminder items.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/reminderitemactions
import "encoding/xml"

type ReminderItemActions struct {
	XMLName xml.Name

	// The ReminderItemAction element specifies the action for a reminder item.
	ReminderItemAction *ReminderItemAction `xml:"ReminderItemAction"`
}

func (R *ReminderItemActions) SetForMarshal() {
	R.XMLName.Local = "m:ReminderItemActions"
}

func (R *ReminderItemActions) GetSchema() *Schema {
	return &SchemaMessages
}
