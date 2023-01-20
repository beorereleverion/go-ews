package elements

// The Time element describes the time when the time changes between standard time and daylight saving time.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/time-timechangetype
import "time"

import "encoding/xml"

type TimeTimeChangeType struct {
	XMLName xml.Name
	TEXT    time.Time `xml:",chardata"`
}

func (T *TimeTimeChangeType) SetForMarshal() {
	T.XMLName.Local = "t:Time"
}

func (T *TimeTimeChangeType) GetSchema() *Schema {
	return &SchemaTypes
}
