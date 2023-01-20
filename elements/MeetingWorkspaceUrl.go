package elements

// The MeetingWorkspaceUrl element contains the URL for the meeting workspace that is included in the calendar item. A meeting workspace is a shared Web site for planning the meeting and tracking the results.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/meetingworkspaceurl
import "encoding/xml"

type MeetingWorkspaceUrl struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (M *MeetingWorkspaceUrl) SetForMarshal() {
	M.XMLName.Local = "t:MeetingWorkspaceUrl"
}

func (M *MeetingWorkspaceUrl) GetSchema() *Schema {
	return &SchemaTypes
}
