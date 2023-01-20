package elements

// The RecurringDayTransition element represents a time zone transition that occurs on the same day each year.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/recurringdaytransition
import "encoding/xml"

type RecurringDayTransition struct {
	XMLName xml.Name

	// The DayOfWeek element represents the day of the week on which the time zone transition occurs.
	DayOfWeek *DayOfWeekTimeZone `xml:"DayOfWeek"`
	// The Month element represents the month in which the time zone transition occurs.
	Month *MonthTimeZoneTransition `xml:"Month"`
	// The Occurrence element represents the occurrence of the day of the week in the month that the time zone transition occurs.
	Occurrence *OccurrenceTimeZoneTransition `xml:"Occurrence"`
	// The Period element defines the name, time offset, and unique identifier for a specific stage of the time zone.
	Period *Period `xml:"Period"`
	// The TimeOffset element represents the time offset from Coordinated Universal Time (UTC) for the time zone transition.
	TimeOffset *TimeOffset `xml:"TimeOffset"`
	// The To element specifies the target of the time zone transition. The target is either a time zone period or a group of time zone transitions.
	To *To `xml:"To"`
	// The TransitionsGroup element represents an array of time zone transitions.
	TransitionsGroup *TransitionsGroup `xml:"TransitionsGroup"`
}

func (R *RecurringDayTransition) SetForMarshal() {
	R.XMLName.Local = "t:RecurringDayTransition"
}

func (R *RecurringDayTransition) GetSchema() *Schema {
	return &SchemaTypes
}
