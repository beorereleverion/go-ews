package elements

// The PerformReminderAction element specifies a request to perform a reminder action.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/performreminderaction
import "encoding/xml"

type PerformReminderAction struct {
	XMLName xml.Name

	// The ReminderItemActions element specifies the actions for reminder items.
	ReminderItemActions *ReminderItemActions `xml:"ReminderItemActions"`
}

func (P *PerformReminderAction) SetForMarshal() {
	P.XMLName.Local = "m:PerformReminderAction"
}

func (P *PerformReminderAction) GetSchema() *Schema {
	return &SchemaMessages
}
