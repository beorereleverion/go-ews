package elements

// The CreationTime element specifies when the persona was created.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/creationtime
import "time"

import "encoding/xml"

type CreationTime struct {
	XMLName xml.Name
	TEXT    time.Time `xml:",chardata"`
}

func (C *CreationTime) SetForMarshal() {
	C.XMLName.Local = "t:CreationTime"
}

func (C *CreationTime) GetSchema() *Schema {
	return &SchemaTypes
}
