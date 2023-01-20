package elements

// The RecurringDateTransition element represents a time zone transition that occurs on a specific date each year.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/recurringdatetransition
import "encoding/xml"

type RecurringDateTransition struct {
	XMLName xml.Name

	// The Day element represents the day of the month on which the time zone transition occurs.
	Day *Day `xml:"Day"`
	// The Month element represents the month in which the time zone transition occurs.
	Month *MonthTimeZoneTransition `xml:"Month"`
	// The Period element defines the name, time offset, and unique identifier for a specific stage of the time zone.
	Period *Period `xml:"Period"`
	// The TimeOffset element represents the time offset from Coordinated Universal Time (UTC) for the time zone transition.
	TimeOffset *TimeOffset `xml:"TimeOffset"`
	// The To element specifies the target of the time zone transition. The target is either a time zone period or a group of time zone transitions.
	To *To `xml:"To"`
	// The TransitionsGroup element represents an array of time zone transitions.
	TransitionsGroup *TransitionsGroup `xml:"TransitionsGroup"`
}

func (R *RecurringDateTransition) SetForMarshal() {
	R.XMLName.Local = "t:RecurringDateTransition"
}

func (R *RecurringDateTransition) GetSchema() *Schema {
	return &SchemaTypes
}
