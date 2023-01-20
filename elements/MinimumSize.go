package elements

// The MinimumSize element represents the minimum size that a message must be in order for the condition or exception to apply.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/minimumsize
import "encoding/xml"

type MinimumSize struct {
	XMLName xml.Name
	TEXT    int64 `xml:",chardata"`
}

func (M *MinimumSize) SetForMarshal() {
	M.XMLName.Local = "m:MinimumSize"
}

func (M *MinimumSize) GetSchema() *Schema {
	return &SchemaMessages
}
