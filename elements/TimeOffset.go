package elements

// The TimeOffset element represents the time offset from Coordinated Universal Time (UTC) for the time zone transition.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/timeoffset
import "encoding/xml"

type TimeOffset struct {
	XMLName xml.Name
	TEXT    int64 `xml:",chardata"`
}

func (T *TimeOffset) SetForMarshal() {
	T.XMLName.Local = "t:TimeOffset"
}

func (T *TimeOffset) GetSchema() *Schema {
	return &SchemaTypes
}
