package elements

// The Duration element specifies the duration that the out of office (OOF) status is enabled if the OofState element is set to Scheduled.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/duration-useroofsettings
type DurationUserOofSettings struct {
	// The EndTime element represents the end of a time span.
	EndTime *EndTime `xml:"t:EndTime"`
	// The StartTime element represents the start of a time span.
	StartTime *StartTime `xml:"t:StartTime"`
}
