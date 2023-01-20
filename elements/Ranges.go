package elements

// The Ranges element specifies an array of recurrence ranges.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/ranges
import "encoding/xml"

type Ranges struct {
	XMLName xml.Name

	// The Range element specifies a range of calendar item occurrences for a repeating calendar item.
	Range *Range `xml:"Range"`
}

func (R *Ranges) SetForMarshal() {
	R.XMLName.Local = "t:Ranges"
}

func (R *Ranges) GetSchema() *Schema {
	return &SchemaTypes
}
