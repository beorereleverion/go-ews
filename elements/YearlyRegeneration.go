package elements

// The YearlyRegeneration element describes the frequency, in years, in which a task is regenerated.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/yearlyregeneration
import "encoding/xml"

type YearlyRegeneration struct {
	XMLName xml.Name

	// The Interval element defines the interval between two consecutive recurring items.
	Interval *Interval `xml:"Interval"`
}

func (Y *YearlyRegeneration) SetForMarshal() {
	Y.XMLName.Local = "t:YearlyRegeneration"
}

func (Y *YearlyRegeneration) GetSchema() *Schema {
	return &SchemaTypes
}
