package elements

// The WorkCitySortKey element contains the sort key for the work city values associated with a persona.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/workcitysortkey
import "encoding/xml"

type WorkCitySortKey struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (W *WorkCitySortKey) SetForMarshal() {
	W.XMLName.Local = "t:WorkCitySortKey"
}

func (W *WorkCitySortKey) GetSchema() *Schema {
	return &SchemaTypes
}
