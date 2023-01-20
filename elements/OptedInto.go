package elements

// The OptedInto element specifies a Boolean value that indicates whether the user opted in to the retention policy.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/optedinto
import "encoding/xml"

type OptedInto struct {
	XMLName xml.Name
	TEXT    bool `xml:",chardata"`
}

const (
	// false
	OptedIntofalse bool = false
	// true
	OptedIntotrue bool = true
)

func (O *OptedInto) SetForMarshal() {
	O.XMLName.Local = "t:OptedInto"
}

func (O *OptedInto) GetSchema() *Schema {
	return &SchemaTypes
}
