package elements

// The LastModifiedTime element indicates when an item was last modified. This element is read-only.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/lastmodifiedtime
import "time"

import "encoding/xml"

type LastModifiedTime struct {
	XMLName xml.Name
	TEXT    time.Time `xml:",chardata"`
}

func (L *LastModifiedTime) SetForMarshal() {
	L.XMLName.Local = "t:LastModifiedTime"
}

func (L *LastModifiedTime) GetSchema() *Schema {
	return &SchemaTypes
}
