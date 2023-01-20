package elements

// The MeetingRequestType element describes the type of the meeting request.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/meetingrequesttype
import "encoding/xml"

type MeetingRequestType struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

const (
	// Identifies the meeting request as a full update to an existing request. A full update has updated time and informational content.
	MeetingRequestTypeFullUpdate string = `FullUpdate`
	// Identifies the meeting request as only containing updated informational content.
	MeetingRequestTypeInformationalUpdate string = `InformationalUpdate`
	// Identifies the meeting request as a new meeting request.
	MeetingRequestTypeNewMeetingRequest string = `NewMeetingRequest`
	// Indicates that the meeting request type is not defined.
	MeetingRequestTypeNone string = `None`
	// Identifies the meeting request as outdated.
	MeetingRequestTypeOutdated string = `Outdated`
	// Indicates that the meeting request belongs to a principal who has forwarded meeting messages to a delegate and has his copies marked as informational.
	MeetingRequestTypePrincipalWantsCopy string = `PrincipalWantsCopy`
	// Identifies the meeting request as a silent update to an existing meeting.
	MeetingRequestTypeSilentUpdate string = `SilentUpdate`
)

func (M *MeetingRequestType) SetForMarshal() {
	M.XMLName.Local = "t:MeetingRequestType"
}

func (M *MeetingRequestType) GetSchema() *Schema {
	return &SchemaTypes
}
