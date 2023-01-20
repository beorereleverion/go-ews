package elements

// The AbsoluteDateTransition element represents a time zone transition that occurs on a specific date and at a specific time.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/absolutedatetransition
import "encoding/xml"

type AbsoluteDateTransition struct {
	XMLName xml.Name

	// The DateTime element represents the date and time at which the time zone transition occurs.
	DateTime *DateTime `xml:"DateTime"`
	// The Period element defines the name, time offset, and unique identifier for a specific stage of the time zone.
	Period *Period `xml:"Period"`
	// The To element specifies the target of the time zone transition. The target is either a time zone period or a group of time zone transitions.
	To *To `xml:"To"`
	// The TransitionsGroup element represents an array of time zone transitions.
	TransitionsGroup *TransitionsGroup `xml:"TransitionsGroup"`
}

func (A *AbsoluteDateTransition) SetForMarshal() {
	A.XMLName.Local = "t:AbsoluteDateTransition"
}

func (A *AbsoluteDateTransition) GetSchema() *Schema {
	return &SchemaTypes
}
