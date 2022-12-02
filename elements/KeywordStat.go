package elements

// The KeywordStat element specifies keyword statistic information.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/keywordstat
type KeywordStat struct {
	// The ItemHits element identifies how many times a keyword was found.
	ItemHits *ItemHits `xml:"ItemHits"`
	// The Keyword element specifies a single keyword.
	Keyword *Keyword `xml:"Keyword"`
	// The Size element specifies the total size of one or more mailbox items.
	Size *Sizelong `xml:"t:Size"`
}
