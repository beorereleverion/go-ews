package elements

// The MeetingRequestWasSent element indicates whether a meeting request has been sent to requested attendees.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/meetingrequestwassent
import "encoding/xml"

type MeetingRequestWasSent struct {
	XMLName xml.Name
	TEXT    bool `xml:",chardata"`
}

func (M *MeetingRequestWasSent) SetForMarshal() {
	M.XMLName.Local = "t:MeetingRequestWasSent"
}

func (M *MeetingRequestWasSent) GetSchema() *Schema {
	return &SchemaTypes
}
