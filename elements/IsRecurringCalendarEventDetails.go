package elements

// The IsRecurring element indicates whether the calendar event is an instance of a recurring calendar item or a single calendar item.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/isrecurring-calendareventdetails
import "encoding/xml"

type IsRecurringCalendarEventDetails struct {
	XMLName xml.Name
	TEXT    bool `xml:",chardata"`
}

const (
	// false
	IsRecurringCalendarEventDetailsfalse bool = false
	// true
	IsRecurringCalendarEventDetailstrue bool = true
)

func (I *IsRecurringCalendarEventDetails) SetForMarshal() {
	I.XMLName.Local = "t:IsRecurring"
}

func (I *IsRecurringCalendarEventDetails) GetSchema() *Schema {
	return &SchemaTypes
}
