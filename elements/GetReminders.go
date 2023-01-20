package elements

// The GetReminders element specifies a request to get reminders.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/getreminders
import "encoding/xml"

type GetReminders struct {
	XMLName xml.Name

	// The BeginTime element specifies the beginning of the time span to query for reminders.
	BeginTime *BeginTime `xml:"BeginTime"`
	// The EndTime element represents the end of the time span to query for reminders.
	EndTime *EndTimeReminderMessageDataType `xml:"EndTime"`
	// The MaxItems element specifies the maximum number of items to return in the request.
	MaxItems *MaxItems `xml:"MaxItems"`
	// The ReminderType element specifies the type of reminders to return.
	ReminderType *ReminderType `xml:"ReminderType"`
}

func (G *GetReminders) SetForMarshal() {
	G.XMLName.Local = "m:GetReminders"
}

func (G *GetReminders) GetSchema() *Schema {
	return &SchemaMessages
}
