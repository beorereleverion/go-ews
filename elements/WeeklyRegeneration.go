package elements

// The WeeklyRegeneration element describes the frequency, in weeks, in which a task is regenerated.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/weeklyregeneration
import "encoding/xml"

type WeeklyRegeneration struct {
	XMLName xml.Name

	// The Interval element defines the interval between two consecutive recurring items.
	Interval *Interval `xml:"Interval"`
}

func (W *WeeklyRegeneration) SetForMarshal() {
	W.XMLName.Local = "t:WeeklyRegeneration"
}

func (W *WeeklyRegeneration) GetSchema() *Schema {
	return &SchemaTypes
}
