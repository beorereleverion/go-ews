package elements

// The MailboxStatisticsSearchResult element contains the results of a keyword search.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/mailboxstatisticssearchresult
type MailboxStatisticsSearchResult struct {
	// The KeywordStatisticsSearchResult element contains a single keyword search result.
	KeywordStatisticsSearchResult *KeywordStatisticsSearchResult `xml:"t:KeywordStatisticsSearchResult"`
	// The UserMailbox element identifies a user mailbox.
	UserMailbox *UserMailbox `xml:"t:UserMailbox"`
}
