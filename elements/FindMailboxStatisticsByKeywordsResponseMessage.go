package elements

// The FindMailboxStatisticsByKeywordsResponseMessage element specifies the response message for a FindMailboxStatisticsByKeywords request.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/findmailboxstatisticsbykeywordsresponsemessage
type FindMailboxStatisticsByKeywordsResponseMessage struct {
	// The DescriptiveLinkKey element is currently unused and is reserved for future use. It contains a value of 0.
	DescriptiveLinkKey *DescriptiveLinkKey `xml:"m:DescriptiveLinkKey"`
	// The MailboxStatisticsSearchResult element contains the results of a keyword search.
	MailboxStatisticsSearchResult *MailboxStatisticsSearchResult `xml:"MailboxStatisticsSearchResult"`
	// The MessageText element provides a text description of the status of the response.
	MessageText *MessageText `xml:"m:MessageText"`
	// The MessageXml element provides additional error response information.
	MessageXml *MessageXml `xml:"m:MessageXml"`
	// The ResponseCode element provides status information about the request.
	ResponseCode *ResponseCode `xml:"m:ResponseCode"`
	// Specifies the response class.
	ResponseClass *string `xml:"ResponseClass,attr"`
}
