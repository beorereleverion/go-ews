package elements

// The CurrentMeetingTime element represents the start time of a meeting that you want to update with a meeting time proposed by a meeting attendee.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/currentmeetingtime
import "time"

import "encoding/xml"

type CurrentMeetingTime struct {
	XMLName xml.Name
	TEXT    time.Time `xml:",chardata"`
}

func (C *CurrentMeetingTime) SetForMarshal() {
	C.XMLName.Local = "t:CurrentMeetingTime"
}

func (C *CurrentMeetingTime) GetSchema() *Schema {
	return &SchemaTypes
}
