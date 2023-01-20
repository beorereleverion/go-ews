package elements

// The IsAutomaticForward element indicates whether incoming messages must be automatic forwards in order for the condition or exception to apply.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/isautomaticforward
import "encoding/xml"

type IsAutomaticForward struct {
	XMLName xml.Name
	TEXT    bool `xml:",chardata"`
}

const (
	// false
	IsAutomaticForwardfalse bool = false
	// true
	IsAutomaticForwardtrue bool = true
)

func (I *IsAutomaticForward) SetForMarshal() {
	I.XMLName.Local = "m:IsAutomaticForward"
}

func (I *IsAutomaticForward) GetSchema() *Schema {
	return &SchemaMessages
}
