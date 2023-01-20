package elements

// The Month element represents the transition month of the year to and from standard time and daylight saving time.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/month
import "encoding/xml"

type Month struct {
	XMLName xml.Name
	TEXT    int64 `xml:",chardata"`
}

func (M *Month) SetForMarshal() {
	M.XMLName.Local = "t:Month"
}

func (M *Month) GetSchema() *Schema {
	return &SchemaTypes
}
