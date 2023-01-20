package elements

// The MeetingDurationInMinutes element specifies the duration of the meeting to be suggested.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/meetingdurationinminutes
import "encoding/xml"

type MeetingDurationInMinutes struct {
	XMLName xml.Name
	TEXT    int64 `xml:",chardata"`
}

func (M *MeetingDurationInMinutes) SetForMarshal() {
	M.XMLName.Local = "t:MeetingDurationInMinutes"
}

func (M *MeetingDurationInMinutes) GetSchema() *Schema {
	return &SchemaTypes
}
