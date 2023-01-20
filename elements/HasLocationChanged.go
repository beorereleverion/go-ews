package elements

// The HasLocationChanged element specifies whether the location property of a meeting has changed.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/haslocationchanged
import "encoding/xml"

type HasLocationChanged struct {
	XMLName xml.Name
	TEXT    bool `xml:",chardata"`
}

const (
	// false
	HasLocationChangedfalse bool = false
	// true
	HasLocationChangedtrue bool = true
)

func (H *HasLocationChanged) SetForMarshal() {
	H.XMLName.Local = "t:HasLocationChanged"
}

func (H *HasLocationChanged) GetSchema() *Schema {
	return &SchemaTypes
}
