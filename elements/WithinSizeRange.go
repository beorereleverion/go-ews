package elements

// The WithinSizeRange element specifies the minimum and maximum sizes that incoming messages must be in order for the condition or exception to apply.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/withinsizerange
import "encoding/xml"

type WithinSizeRange struct {
	XMLName xml.Name

	// The MaximumSize element represents the maximum size that a message must be in order for the condition or exception to apply.
	MaximumSize *MaximumSize `xml:"MaximumSize"`
	// The MinimumSize element represents the minimum size that a message must be in order for the condition or exception to apply.
	MinimumSize *MinimumSize `xml:"MinimumSize"`
}

func (W *WithinSizeRange) SetForMarshal() {
	W.XMLName.Local = "m:WithinSizeRange"
}

func (W *WithinSizeRange) GetSchema() *Schema {
	return &SchemaMessages
}
