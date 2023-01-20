package elements

// The Time element represents the transition time of day to and from standard time and daylight saving time.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/time
import "encoding/xml"

type Time struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (T *Time) SetForMarshal() {
	T.XMLName.Local = "t:Time"
}

func (T *Time) GetSchema() *Schema {
	return &SchemaTypes
}
