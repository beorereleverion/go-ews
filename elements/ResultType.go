package elements

// The ResultType element contains the type of search to perform. The type of search can be statistics only or preview only.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/resulttype
type ResultType string

const (
	// PreviewOnly
	ResultTypePreviewOnly ResultType = `PreviewOnly`
	// StatisticsOnly
	ResultTypeStatisticsOnly ResultType = `StatisticsOnly`
)
