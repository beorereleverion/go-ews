package elements

// The OutOfOffice element represents the response message and a duration time for sending the response message.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/outofoffice
type OutOfOffice struct {
	// The Duration element specifies the duration that the out of office (OOF) status is enabled if the OofState element is set to Scheduled.
	Duration *DurationUserOofSettings `xml:"t:Duration"`
	// The OofState element is used to get or set the user's Out of Office (OOF) state.
	OofState *OofState `xml:"t:OofState"`
	// The ReplyBody element contains an Out of Office (OOF) message and the language used for the message.
	ReplyBody *ReplyBody `xml:"t:ReplyBody"`
}
