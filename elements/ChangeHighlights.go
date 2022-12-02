package elements

// The ChangeHighlights element specifies what has changed between two versions of a meeting request message.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/changehighlights
type ChangeHighlights struct {
	// The End element represents the end of a duration.
	End *End `xml:"t:End"`
	// The HasEndTimeChanged element specifies whether the end time for a meeting has changed.
	HasEndTimeChanged *HasEndTimeChanged `xml:"t:HasEndTimeChanged"`
	// The HasLocationChanged element specifies whether the location property of a meeting has changed.
	HasLocationChanged *HasLocationChanged `xml:"t:HasLocationChanged"`
	// The HasStartTimeChanged element specifies whether the start time for a meeting has changed.
	HasStartTimeChanged *HasStartTimeChanged `xml:"t:HasStartTimeChanged"`
	// The Location element represents the location of a meeting, appointment, or persona.
	Location *Location `xml:"t:Location"`
	// The Start element represents the start of a duration.
	Start *Start `xml:"t:Start"`
}
