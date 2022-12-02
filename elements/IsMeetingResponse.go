package elements

// The IsMeetngResponsequest element indicates whether incoming messages must be a meeting response in order for the condition or exception to apply.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/ismeetingresponse
type IsMeetingResponse bool

const (
	// false
	IsMeetingResponsefalse IsMeetingResponse = false
	// true
	IsMeetingResponsetrue IsMeetingResponse = true
)
