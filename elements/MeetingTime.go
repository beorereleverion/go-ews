package elements

// The MeetingTime element represents a suggested meeting time.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/meetingtime
import "time"

import "encoding/xml"

type MeetingTime struct {
	XMLName xml.Name
	TEXT    time.Time `xml:",chardata"`
}

func (M *MeetingTime) SetForMarshal() {
	M.XMLName.Local = "t:MeetingTime"
}

func (M *MeetingTime) GetSchema() *Schema {
	return &SchemaTypes
}
