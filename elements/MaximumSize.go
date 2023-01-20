package elements

// The MaximumSize element represents the maximum size that a message must be in order for the condition or exception to apply.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/maximumsize
import "encoding/xml"

type MaximumSize struct {
	XMLName xml.Name
	TEXT    int64 `xml:",chardata"`
}

func (M *MaximumSize) SetForMarshal() {
	M.XMLName.Local = "m:MaximumSize"
}

func (M *MaximumSize) GetSchema() *Schema {
	return &SchemaMessages
}
