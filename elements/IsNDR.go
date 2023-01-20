package elements

// The IsNDR element indicates whether incoming messages must be non-delivery reports (NDRs) in order for the condition or exception to apply.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/isndr
import "encoding/xml"

type IsNDR struct {
	XMLName xml.Name
	TEXT    bool `xml:",chardata"`
}

const (
	// false
	IsNDRfalse bool = false
	// true
	IsNDRtrue bool = true
)

func (I *IsNDR) SetForMarshal() {
	I.XMLName.Local = "m:IsNDR"
}

func (I *IsNDR) GetSchema() *Schema {
	return &SchemaMessages
}
