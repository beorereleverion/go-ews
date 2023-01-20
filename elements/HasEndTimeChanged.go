package elements

// The HasEndTimeChanged element specifies whether the end time for a meeting has changed.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/hasendtimechanged
import "encoding/xml"

type HasEndTimeChanged struct {
	XMLName xml.Name
	TEXT    bool `xml:",chardata"`
}

const (
	// false
	HasEndTimeChangedfalse bool = false
	// true
	HasEndTimeChangedtrue bool = true
)

func (H *HasEndTimeChanged) SetForMarshal() {
	H.XMLName.Local = "t:HasEndTimeChanged"
}

func (H *HasEndTimeChanged) GetSchema() *Schema {
	return &SchemaTypes
}
