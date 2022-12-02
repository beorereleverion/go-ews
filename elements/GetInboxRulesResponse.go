package elements

// The GetInboxRulesResponse element defines a response to a GetInboxRules operation request.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/getinboxrulesresponse
type GetInboxRulesResponse struct {
	// The DescriptiveLinkKey element is currently unused and is reserved for future use. It contains a value of 0.
	DescriptiveLinkKey *DescriptiveLinkKey `xml:"m:DescriptiveLinkKey"`
	// The InboxRules element represents an array of rules in the user's mailbox.
	InboxRules *InboxRules `xml:"m:InboxRules"`
	// The MessageText element provides a text description of the status of the response.
	MessageText *MessageText `xml:"m:MessageText"`
	// The MessageXml element provides additional error response information.
	MessageXml *MessageXml `xml:"m:MessageXml"`
	// The OutlookRuleBlobExists element indicates whether a Microsoft Outlook rule blob exists in the user's mailbox.
	OutlookRuleBlobExists *OutlookRuleBlobExists `xml:"m:OutlookRuleBlobExists"`
	// The ResponseCode element provides status information about the request.
	ResponseCode *ResponseCode `xml:"m:ResponseCode"`
	// Describes the status of a GetInboxRules operation response. The following values are valid for this attribute:  - Success  - Warning  - Error
	ResponseClass *string `xml:"ResponseClass,attr"`
}
