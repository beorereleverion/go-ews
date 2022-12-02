package elements

// The DistinguishedGroupBy element provides standard groupings for FindItem queries.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/distinguishedgroupby
type DistinguishedGroupBy struct {
	// The StandardGroupBy element represents the standard grouping and aggregating mechanisms for a grouped FindItem operation.
	StandardGroupBy *StandardGroupBy `xml:"t:StandardGroupBy"`
}
