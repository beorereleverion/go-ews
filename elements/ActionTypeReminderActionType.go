package elements

// The ActionType element specifies the action to take on the reminder.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/actiontype-reminderactiontype
import "encoding/xml"

type ActionTypeReminderActionType struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

const (
	// Dismiss
	ActionTypeReminderActionTypeDismiss string = `Dismiss`
	// Snooze
	ActionTypeReminderActionTypeSnooze string = `Snooze`
)

func (A *ActionTypeReminderActionType) SetForMarshal() {
	A.XMLName.Local = "t:ActionType"
}

func (A *ActionTypeReminderActionType) GetSchema() *Schema {
	return &SchemaTypes
}
