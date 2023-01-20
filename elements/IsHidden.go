package elements

// The IsHidden element contains a Boolean value that indicates whether the underlying contact should be hidden or displayed as part of the persona.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/ishidden
import "encoding/xml"

type IsHidden struct {
	XMLName xml.Name
	TEXT    bool `xml:",chardata"`
}

const (
	// false
	IsHiddenfalse bool = false
	// true
	IsHiddentrue bool = true
)

func (I *IsHidden) SetForMarshal() {
	I.XMLName.Local = "t:IsHidden"
}

func (I *IsHidden) GetSchema() *Schema {
	return &SchemaTypes
}
