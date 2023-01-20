package elements

// The HasStartTimeChanged element specifies whether the start time for a meeting has changed.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/hasstarttimechanged
import "encoding/xml"

type HasStartTimeChanged struct {
	XMLName xml.Name
	TEXT    bool `xml:",chardata"`
}

const (
	// false
	HasStartTimeChangedfalse bool = false
	// true
	HasStartTimeChangedtrue bool = true
)

func (H *HasStartTimeChanged) SetForMarshal() {
	H.XMLName.Local = "t:HasStartTimeChanged"
}

func (H *HasStartTimeChanged) GetSchema() *Schema {
	return &SchemaTypes
}
