package elements

// The NonIndexableItemStatistics element contains an array of statistics for items that could not be indexed.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/nonindexableitemstatistics
type NonIndexableItemStatistics struct {
	// The NonIndexableItemStatistic element contains a single statistic for an item that could not be indexed
	NonIndexableItemStatistic *NonIndexableItemStatistic `xml:"m:NonIndexableItemStatistic"`
}
