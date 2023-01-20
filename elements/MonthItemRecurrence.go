package elements

// The Month element describes the month when a yearly recurring item occurs.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/month-item-recurrence
import "encoding/xml"

type MonthItemRecurrence struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (M *MonthItemRecurrence) SetForMarshal() {
	M.XMLName.Local = "t:Month"
}

func (M *MonthItemRecurrence) GetSchema() *Schema {
	return &SchemaTypes
}
