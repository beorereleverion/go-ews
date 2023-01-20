package elements

// The IsWorkTime element represents whether the suggested meeting time occurs during scheduled work hours.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/isworktime
import "encoding/xml"

type IsWorkTime struct {
	XMLName xml.Name
	TEXT    bool `xml:",chardata"`
}

const (
	// false
	IsWorkTimefalse bool = false
	// true
	IsWorkTimetrue bool = true
)

func (I *IsWorkTime) SetForMarshal() {
	I.XMLName.Local = "t:IsWorkTime"
}

func (I *IsWorkTime) GetSchema() *Schema {
	return &SchemaTypes
}
