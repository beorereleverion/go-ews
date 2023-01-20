package elements

// The DateTimePrecision element specifies the precision for returned date/time values.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/datetimeprecision
import "encoding/xml"

type DateTimePrecision struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (D *DateTimePrecision) SetForMarshal() {
	D.XMLName.Local = "t:DateTimePrecision"
}

func (D *DateTimePrecision) GetSchema() *Schema {
	return &SchemaTypes
}
