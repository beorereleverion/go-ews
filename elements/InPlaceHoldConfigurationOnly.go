package elements

// The InPlaceHoldConfigurationOnly element specifies whether to include the in-place hold configuration.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/inplaceholdconfigurationonly
import "encoding/xml"

type InPlaceHoldConfigurationOnly struct {
	XMLName xml.Name
	TEXT    bool `xml:",chardata"`
}

const (
	// false
	InPlaceHoldConfigurationOnlyfalse bool = false
	// true
	InPlaceHoldConfigurationOnlytrue bool = true
)

func (I *InPlaceHoldConfigurationOnly) SetForMarshal() {
	I.XMLName.Local = "m:InPlaceHoldConfigurationOnly"
}

func (I *InPlaceHoldConfigurationOnly) GetSchema() *Schema {
	return &SchemaMessages
}
