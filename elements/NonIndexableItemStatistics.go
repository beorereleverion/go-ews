package elements

// The NonIndexableItemStatistics element contains an array of statistics for items that could not be indexed.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/nonindexableitemstatistics
import "encoding/xml"

type NonIndexableItemStatistics struct {
	XMLName xml.Name

	// The NonIndexableItemStatistic element contains a single statistic for an item that could not be indexed
	NonIndexableItemStatistic *NonIndexableItemStatistic `xml:"NonIndexableItemStatistic"`
}

func (N *NonIndexableItemStatistics) SetForMarshal() {
	N.XMLName.Local = "m:NonIndexableItemStatistics"
}

func (N *NonIndexableItemStatistics) GetSchema() *Schema {
	return &SchemaMessages
}
