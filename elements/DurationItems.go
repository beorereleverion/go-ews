package elements

// The Duration element represents the duration of a calendar item.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/duration-items
import "encoding/xml"

type DurationItems struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (D *DurationItems) SetForMarshal() {
	D.XMLName.Local = "t:Duration"
}

func (D *DurationItems) GetSchema() *Schema {
	return &SchemaTypes
}
