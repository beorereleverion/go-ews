package elements

// The WithinDateRange element specifies the date range within which incoming messages have to have been received in order for the condition or exception to apply.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/withindaterange
import "encoding/xml"

type WithinDateRange struct {
	XMLName xml.Name

	// The EndDateTime element specifies the end date and time for a rule or a search.
	EndDateTime *EndDateTime `xml:"EndDateTime"`
	// The StartDateTime element specifies the start date and time for a rule or a search.
	StartDateTime *StartDateTime `xml:"StartDateTime"`
}

func (W *WithinDateRange) SetForMarshal() {
	W.XMLName.Local = "m:WithinDateRange"
}

func (W *WithinDateRange) GetSchema() *Schema {
	return &SchemaMessages
}
