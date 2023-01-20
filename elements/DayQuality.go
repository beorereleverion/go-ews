package elements

// The DayQuality element represents the quality of the day for containing quality suggested meeting times.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/dayquality
import "encoding/xml"

type DayQuality struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

const (
	// Excellent
	DayQualityExcellent string = `Excellent`
	// Fair
	DayQualityFair string = `Fair`
	// Good
	DayQualityGood string = `Good`
	// Poor
	DayQualityPoor string = `Poor`
)

func (D *DayQuality) SetForMarshal() {
	D.XMLName.Local = "t:DayQuality"
}

func (D *DayQuality) GetSchema() *Schema {
	return &SchemaTypes
}
