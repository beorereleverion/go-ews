package elements

// The Date element represents the date and time at which the event occurred.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/date-messagetracking
import "time"

import "encoding/xml"

type DateMessageTracking struct {
	XMLName xml.Name
	TEXT    time.Time `xml:",chardata"`
}

func (D *DateMessageTracking) SetForMarshal() {
	D.XMLName.Local = "t:Date"
}

func (D *DateMessageTracking) GetSchema() *Schema {
	return &SchemaTypes
}
