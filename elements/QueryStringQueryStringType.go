package elements

// The QueryString element contains a mailbox query string based on Advanced Query Syntax (AQS).
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/querystring-querystringtype
type QueryStringQueryStringType struct {
	// Indicates that the cache should be reset.
	ResetCache *string `xml:"ResetCache,attr"`
	// Indicates that deleted items should be returned.
	ReturnDeletedItems *string `xml:"ReturnDeletedItems,attr"`
	// Indicates that highlighted terms should be returned.
	ReturnHighlightTerms *string `xml:"ReturnHighlightTerms,attr"`
}
