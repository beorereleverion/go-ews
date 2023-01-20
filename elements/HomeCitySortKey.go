package elements

// The HomeCitySortKey element represents the sort key for the home city.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/homecitysortkey
import "encoding/xml"

type HomeCitySortKey struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (H *HomeCitySortKey) SetForMarshal() {
	H.XMLName.Local = "t:HomeCitySortKey"
}

func (H *HomeCitySortKey) GetSchema() *Schema {
	return &SchemaTypes
}
