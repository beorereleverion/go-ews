package elements

// The ChangeHighlights element specifies what has changed between two versions of a meeting request message.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/changehighlights
import "encoding/xml"

type ChangeHighlights struct {
	XMLName xml.Name

	// The End element represents the end of a duration.
	End *End `xml:"End"`
	// The HasEndTimeChanged element specifies whether the end time for a meeting has changed.
	HasEndTimeChanged *HasEndTimeChanged `xml:"HasEndTimeChanged"`
	// The HasLocationChanged element specifies whether the location property of a meeting has changed.
	HasLocationChanged *HasLocationChanged `xml:"HasLocationChanged"`
	// The HasStartTimeChanged element specifies whether the start time for a meeting has changed.
	HasStartTimeChanged *HasStartTimeChanged `xml:"HasStartTimeChanged"`
	// The Location element represents the location of a meeting, appointment, or persona.
	Location *Location `xml:"Location"`
	// The Start element represents the start of a duration.
	Start *Start `xml:"Start"`
}

func (C *ChangeHighlights) SetForMarshal() {
	C.XMLName.Local = "t:ChangeHighlights"
}

func (C *ChangeHighlights) GetSchema() *Schema {
	return &SchemaTypes
}
