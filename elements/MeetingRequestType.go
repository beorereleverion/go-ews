package elements

// The MeetingRequestType element describes the type of the meeting request.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/meetingrequesttype
type MeetingRequestType string

const (
	// Identifies the meeting request as a full update to an existing request. A full update has updated time and informational content.
	MeetingRequestTypeFullUpdate MeetingRequestType = `FullUpdate`
	// Identifies the meeting request as only containing updated informational content.
	MeetingRequestTypeInformationalUpdate MeetingRequestType = `InformationalUpdate`
	// Identifies the meeting request as a new meeting request.
	MeetingRequestTypeNewMeetingRequest MeetingRequestType = `NewMeetingRequest`
	// Indicates that the meeting request type is not defined.
	MeetingRequestTypeNone MeetingRequestType = `None`
	// Identifies the meeting request as outdated.
	MeetingRequestTypeOutdated MeetingRequestType = `Outdated`
	// Indicates that the meeting request belongs to a principal who has forwarded meeting messages to a delegate and has his copies marked as informational.
	MeetingRequestTypePrincipalWantsCopy MeetingRequestType = `PrincipalWantsCopy`
	// Identifies the meeting request as a silent update to an existing meeting.
	MeetingRequestTypeSilentUpdate MeetingRequestType = `SilentUpdate`
)
