package elements

// The ProposeNewTime element specifies a response object that indicates that the meeting attendee can propose a new meeting time.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/proposenewtime
type ProposeNewTime struct {
	// The name of the response object.
	ObjectName *string `xml:"ObjectName,attr"`
}
