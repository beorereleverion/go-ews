package elements

// The MonthlyRegeneration element describes the frequency, in months, of which task is regenerated.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/monthlyregeneration
import "encoding/xml"

type MonthlyRegeneration struct {
	XMLName xml.Name

	// The Interval element defines the interval between two consecutive recurring items.
	Interval *Interval `xml:"Interval"`
}

func (M *MonthlyRegeneration) SetForMarshal() {
	M.XMLName.Local = "t:MonthlyRegeneration"
}

func (M *MonthlyRegeneration) GetSchema() *Schema {
	return &SchemaTypes
}
