package elements

// The ReminderItemAction element specifies the action for a reminder item.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/reminderitemaction
import "encoding/xml"

type ReminderItemAction struct {
	XMLName xml.Name

	// The ActionType element specifies the action to take on the reminder.
	ActionType *ActionTypeReminderActionType `xml:"ActionType"`
	// The ItemId element contains the unique identifier and change key of an item in the Exchange store.
	ItemId *ItemId `xml:"ItemId"`
	// The NewReminderTime element specifies a new time for a reminder.
	NewReminderTime *NewReminderTime `xml:"NewReminderTime"`
}

func (R *ReminderItemAction) SetForMarshal() {
	R.XMLName.Local = "t:ReminderItemAction"
}

func (R *ReminderItemAction) GetSchema() *Schema {
	return &SchemaTypes
}
