package elements

// The MeetingString element specifies the name of the meeting as the result of entity extraction.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/meetingstring
import "encoding/xml"

type MeetingString struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (M *MeetingString) SetForMarshal() {
	M.XMLName.Local = "t:MeetingString"
}

func (M *MeetingString) GetSchema() *Schema {
	return &SchemaTypes
}
