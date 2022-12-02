package elements

// The GetNonIndexableItemStatistics element specifies a request to retrieve nonindexable item statistics.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/getnonindexableitemstatistics
type GetNonIndexableItemStatistics struct {
	// The Mailboxes element specifies an array of mailboxes identified by legacy distinguished name.
	Mailboxes *MailboxesNonEmptyArrayOfLegacyDNsType `xml:"m:Mailboxes"`
}
