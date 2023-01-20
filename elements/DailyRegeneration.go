package elements

// The DailyRegeneration element describes the frequency, in days, in which a task is regenerated.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/dailyregeneration
import "encoding/xml"

type DailyRegeneration struct {
	XMLName xml.Name

	// The Interval element defines the interval between two consecutive recurring items.
	Interval *Interval `xml:"Interval"`
}

func (D *DailyRegeneration) SetForMarshal() {
	D.XMLName.Local = "t:DailyRegeneration"
}

func (D *DailyRegeneration) GetSchema() *Schema {
	return &SchemaTypes
}
