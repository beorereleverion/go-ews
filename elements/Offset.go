package elements

// The Offset element describes the offset from the BaseOffset. Together with the BaseOffset element, the Offset element identifies whether the time is standard or daylight saving time.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/offset
import "encoding/xml"

type Offset struct {
	XMLName xml.Name
	TEXT    int64 `xml:",chardata"`
}

func (O *Offset) SetForMarshal() {
	O.XMLName.Local = "t:Offset"
}

func (O *Offset) GetSchema() *Schema {
	return &SchemaTypes
}
