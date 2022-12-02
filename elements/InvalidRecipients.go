package elements

// The InvalidRecipients element represents the recipients of a folder sharing request that are invalid.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/invalidrecipients
type InvalidRecipients struct {
	// The InvalidRecipient element contains the SMTP address of the invalid recipient and information about why the recipient is invalid.
	InvalidRecipient *InvalidRecipient `xml:"t:InvalidRecipient"`
}
