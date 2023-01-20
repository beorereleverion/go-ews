package elements

// The GetRemindersResponse element specifies the response to a GetReminders request.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/getremindersresponse
import "encoding/xml"

type GetRemindersResponse struct {
	XMLName xml.Name

	// The Reminders element specifies the reminders returned in the response to a GetReminders request.
	Reminders *Reminders `xml:"Reminders"`
}

func (G *GetRemindersResponse) SetForMarshal() {
	G.XMLName.Local = "m:GetRemindersResponse"
}

func (G *GetRemindersResponse) GetSchema() *Schema {
	return &SchemaMessages
}
