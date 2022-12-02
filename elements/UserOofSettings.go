package elements

// The UserOofSettings element specifies the Out of Office (OOF) settings.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/useroofsettings
type UserOofSettings struct {
	// The Duration element specifies the duration that the out of office (OOF) status is enabled if the OofState element is set to Scheduled.
	Duration *DurationUserOofSettings `xml:"t:Duration"`
	// The ExternalAudience element sets or contains a value that determines to whom external Out of Office (OOF) messages are sent.
	ExternalAudience *ExternalAudience `xml:"t:ExternalAudience"`
	// The ExternalReply element contains the out of office (OOF) response that is sent to addresses outside the recipient's domain or trusted domains.
	ExternalReply *ExternalReply `xml:"t:ExternalReply"`
	// The InternalReply element contains the out of office (OOF) response sent to other users in the user's domain or trusted domains.
	InternalReply *InternalReply `xml:"t:InternalReply"`
	// The OofState element is used to get or set the user's Out of Office (OOF) state.
	OofState *OofState `xml:"t:OofState"`
}
