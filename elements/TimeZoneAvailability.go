package elements

// The TimeZone element contains elements that identify time zone information. This element also contains information about the transition between standard time and daylight saving time.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/timezone-availability
import "encoding/xml"

type TimeZoneAvailability struct {
	XMLName xml.Name

	// The DaylightTime element represents an offset from the time relative to Coordinated Universal Time (UTC) that is represented by the Bias (UTC) element in regions where daylight saving time is observed. This element also contains information about when the transition to daylight saving time from standard time occurs.
	DaylightTime *DaylightTime `xml:"DaylightTime"`
	// The StandardTime element represents an offset from the time relative to Coordinated Universal Time (UTC) that is represented by the Bias (UTC) element. This element also contains information about the transition to standard time from daylight saving time in regions where daylight saving time is observed.
	StandardTime *StandardTime `xml:"StandardTime"`
}

func (T *TimeZoneAvailability) SetForMarshal() {
	T.XMLName.Local = "t:TimeZone"
}

func (T *TimeZoneAvailability) GetSchema() *Schema {
	return &SchemaTypes
}
