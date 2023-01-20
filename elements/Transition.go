package elements

// The Transition element represents a time zone transition.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/transition
import "encoding/xml"

type Transition struct {
	XMLName xml.Name

	// The Period element defines the name, time offset, and unique identifier for a specific stage of the time zone.
	Period *Period `xml:"Period"`
	// The To element specifies the target of the time zone transition. The target is either a time zone period or a group of time zone transitions.
	To *To `xml:"To"`
	// The TransitionsGroup element represents an array of time zone transitions.
	TransitionsGroup *TransitionsGroup `xml:"TransitionsGroup"`
}

func (T *Transition) SetForMarshal() {
	T.XMLName.Local = "t:Transition"
}

func (T *Transition) GetSchema() *Schema {
	return &SchemaTypes
}
