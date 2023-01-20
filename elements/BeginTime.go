package elements

// The BeginTime element specifies the beginning of the time span to query for reminders.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/begintime
import "time"

import "encoding/xml"

type BeginTime struct {
	XMLName xml.Name
	TEXT    time.Time `xml:",chardata"`
}

func (B *BeginTime) SetForMarshal() {
	B.XMLName.Local = "m:BeginTime"
}

func (B *BeginTime) GetSchema() *Schema {
	return &SchemaMessages
}
