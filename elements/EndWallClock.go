package elements

// The EndWallClock element specifies the end time of a meeting in the time zone of the location in which the meeting takes place.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/endwallclock
import "time"

import "encoding/xml"

type EndWallClock struct {
	XMLName xml.Name
	TEXT    time.Time `xml:",chardata"`
}

func (E *EndWallClock) SetForMarshal() {
	E.XMLName.Local = "t:EndWallClock"
}

func (E *EndWallClock) GetSchema() *Schema {
	return &SchemaTypes
}
