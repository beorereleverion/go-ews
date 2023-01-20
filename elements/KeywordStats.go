package elements

// The KeywordStats element specifies a list of one or more KeywordStat elements.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/keywordstats
import "encoding/xml"

type KeywordStats struct {
	XMLName xml.Name

	// The KeywordStat element specifies keyword statistic information.
	KeywordStat *KeywordStat `xml:"KeywordStat"`
}
