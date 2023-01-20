package elements

// The StartTime (ReminderMessageDataType) element specifies the starting time of the item that the reminder is for.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/starttime-remindermessagedatatype
import "time"

import "encoding/xml"

type StartTimeReminderMessageDataType struct {
	XMLName xml.Name
	TEXT    time.Time `xml:",chardata"`
}

func (S *StartTimeReminderMessageDataType) SetForMarshal() {
	S.XMLName.Local = "t:StartTime"
}

func (S *StartTimeReminderMessageDataType) GetSchema() *Schema {
	return &SchemaTypes
}
