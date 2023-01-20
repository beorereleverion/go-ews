package elements

// The MaximumBodySize element specifies the maximum size of the item body to return in a response.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/maximumbodysize
import "encoding/xml"

type MaximumBodySize struct {
	XMLName xml.Name
	TEXT    int64 `xml:",chardata"`
}

func (M *MaximumBodySize) SetForMarshal() {
	M.XMLName.Local = "t:MaximumBodySize"
}

func (M *MaximumBodySize) GetSchema() *Schema {
	return &SchemaTypes
}
