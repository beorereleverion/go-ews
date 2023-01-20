package elements

// The LastAttemptTime element contains the time and date at which the last attempt to index the item was made.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/lastattempttime
import "time"

import "encoding/xml"

type LastAttemptTime struct {
	XMLName xml.Name
	TEXT    time.Time `xml:",chardata"`
}

func (L *LastAttemptTime) SetForMarshal() {
	L.XMLName.Local = "t:LastAttemptTime"
}

func (L *LastAttemptTime) GetSchema() *Schema {
	return &SchemaTypes
}
