package elements

// The Occurrence element represents the occurrence of the day of the week in the month that the time zone transition occurs.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/occurrence-time-zone-transition
import "encoding/xml"

type OccurrenceTimeZoneTransition struct {
	XMLName xml.Name
	TEXT    int64 `xml:",chardata"`
}

const (
	// The first occurrence of the specified day of the week from the end of the month.
	OccurrenceTimeZoneTransitionMinusOne int64 = -1
	// The second occurrence of the specified day of the week from the end of the month.
	OccurrenceTimeZoneTransitionMinusTwo int64 = -2
	// The third occurrence of the specified day of the week from the end of the month.
	OccurrenceTimeZoneTransitionMinusTree int64 = -3
	// The fourth occurrence of the specified day of the week from the end of the month.
	OccurrenceTimeZoneTransitionMinusFour int64 = -4
	// The first occurrence of the specified day of the week from the beginning of the month.
	OccurrenceTimeZoneTransitionOne int64 = 1
	// The second occurrence of the specified day of the week from the beginning of the month.
	OccurrenceTimeZoneTransitionTwo int64 = 2
	// The third occurrence of the specified day of the week from the beginning of the month.
	OccurrenceTimeZoneTransitionTree int64 = 3
	// The fourth occurrence of the specified day of the week from the beginning of the month.
	OccurrenceTimeZoneTransitionFour int64 = 4
)

func (O *OccurrenceTimeZoneTransition) SetForMarshal() {
	O.XMLName.Local = "t:Occurrence"
}

func (O *OccurrenceTimeZoneTransition) GetSchema() *Schema {
	return &SchemaTypes
}
