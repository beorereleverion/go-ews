package elements

// The CreatedTime element specifies the time at which the item was created.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/createdtime
import "time"

import "encoding/xml"

type CreatedTime struct {
	XMLName xml.Name
	TEXT    time.Time `xml:",chardata"`
}

func (C *CreatedTime) SetForMarshal() {
	C.XMLName.Local = "t:CreatedTime"
}

func (C *CreatedTime) GetSchema() *Schema {
	return &SchemaTypes
}
