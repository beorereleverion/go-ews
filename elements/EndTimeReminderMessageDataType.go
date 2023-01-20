package elements

// The EndTime element represents the end of the time span to query for reminders.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/endtime-remindermessagedatatype
import "time"

import "encoding/xml"

type EndTimeReminderMessageDataType struct {
	XMLName xml.Name
	TEXT    time.Time `xml:",chardata"`
}

func (E *EndTimeReminderMessageDataType) SetForMarshal() {
	E.XMLName.Local = "t:EndTime"
}

func (E *EndTimeReminderMessageDataType) GetSchema() *Schema {
	return &SchemaTypes
}
