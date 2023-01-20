package elements

// The TimeWindow element identifies the time span queried for the user availability information.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/timewindow
import "encoding/xml"

type TimeWindow struct {
	XMLName xml.Name

	// The EndTime element represents the end of a time span.
	EndTime *EndTime `xml:"EndTime"`
	// The StartTime element represents the start of a time span.
	StartTime *StartTime `xml:"StartTime"`
}

func (T *TimeWindow) SetForMarshal() {
	T.XMLName.Local = "t:TimeWindow"
}

func (T *TimeWindow) GetSchema() *Schema {
	return &SchemaTypes
}
