package elements

// The Transitions element represents an array of time zone transitions.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/transitions
import "encoding/xml"

type Transitions struct {
	XMLName xml.Name

	// The AbsoluteDateTransition element represents a time zone transition that occurs on a specific date and at a specific time.
	AbsoluteDateTransition *AbsoluteDateTransition `xml:"AbsoluteDateTransition"`
	// The RecurringDateTransition element represents a time zone transition that occurs on a specific date each year.
	RecurringDateTransition *RecurringDateTransition `xml:"RecurringDateTransition"`
	// The RecurringDayTransition element represents a time zone transition that occurs on the same day each year.
	RecurringDayTransition *RecurringDayTransition `xml:"RecurringDayTransition"`
	// The Transition element represents a time zone transition.
	Transition *Transition `xml:"Transition"`
	// Represents the unique identifier of the time zone definition.
	Id *string `xml:"Id,attr"`
}

func (T *Transitions) SetForMarshal() {
	T.XMLName.Local = "t:Transitions"
}

func (T *Transitions) GetSchema() *Schema {
	return &SchemaTypes
}
