package elements

// The IncludesLastItemInRange element indicates whether the last item to synchronize has been included in the response.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/includeslastiteminrange
import "encoding/xml"

type IncludesLastItemInRange struct {
	XMLName xml.Name
	TEXT    bool `xml:",chardata"`
}

func (I *IncludesLastItemInRange) SetForMarshal() {
	I.XMLName.Local = "m:IncludesLastItemInRange"
}

func (I *IncludesLastItemInRange) GetSchema() *Schema {
	return &SchemaMessages
}
