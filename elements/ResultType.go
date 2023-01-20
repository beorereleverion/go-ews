package elements

// The ResultType element contains the type of search to perform. The type of search can be statistics only or preview only.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/resulttype
import "encoding/xml"

type ResultType struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

const (
	// PreviewOnly
	ResultTypePreviewOnly string = `PreviewOnly`
	// StatisticsOnly
	ResultTypeStatisticsOnly string = `StatisticsOnly`
)

func (R *ResultType) SetForMarshal() {
	R.XMLName.Local = "m:ResultType"
}

func (R *ResultType) GetSchema() *Schema {
	return &SchemaMessages
}
