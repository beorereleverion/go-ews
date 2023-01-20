package elements

// The WorkingHours element represents the time zone settings and working hours for the requested mailbox user.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/workinghours-ex15websvcsotherref
import "encoding/xml"

type WorkingHours struct {
	XMLName xml.Name

	// The TimeZone element contains elements that identify time zone information. This element also contains information about the transition between standard time and daylight saving time.
	TimeZone *TimeZoneAvailability `xml:"TimeZone"`
	// The WorkingPeriodArray element contains working period information for the mailbox user.
	WorkingPeriodArray *WorkingPeriodArray `xml:"WorkingPeriodArray"`
}

func (W *WorkingHours) SetForMarshal() {
	W.XMLName.Local = "t:WorkingHours"
}

func (W *WorkingHours) GetSchema() *Schema {
	return &SchemaTypes
}
