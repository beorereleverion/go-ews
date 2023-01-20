package elements

// The ToDate element specifies the date that the message was received.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/todate
import "time"

import "encoding/xml"

type ToDate struct {
	XMLName xml.Name
	TEXT    time.Time `xml:",chardata"`
}

func (T *ToDate) SetForMarshal() {
	T.XMLName.Local = "t:ToDate"
}

func (T *ToDate) GetSchema() *Schema {
	return &SchemaTypes
}
