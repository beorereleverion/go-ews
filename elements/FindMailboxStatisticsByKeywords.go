package elements

// The FindMailboxStatisticsByKeywords element specifies a request to search for mailbox statistics by keyword.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/findmailboxstatisticsbykeywords
type FindMailboxStatisticsByKeywords struct {
	// The FromDate element specifies the date that the message was sent.
	FromDate *FromDate `xml:"m:FromDate"`
	// The IncludePersonalArchive element specifies whether to include the personal archive in the search.
	IncludePersonalArchive *IncludePersonalArchive `xml:"m:IncludePersonalArchive"`
	// The IncludeUnsearchableItems element specifies whether to include items that cannot be searched.
	IncludeUnsearchableItems *IncludeUnsearchableItems `xml:"m:IncludeUnsearchableItems"`
	// The Keywords element specifies keywords for a FindMailboxStatisticsByKeywords operation search.
	Keywords *Keywords `xml:"Keywords"`
	// The Language element contains the language used for the search query.
	Language *Language `xml:"m:Language"`
	// The Mailboxes element contains an array of mailboxes.
	Mailboxes *MailboxesArrayOfUserMailboxesType `xml:"m:Mailboxes"`
	// The MessageTypes element specifies an array of messages to search.
	MessageTypes *MessageTypes `xml:"m:MessageTypes"`
	// The Recipients element specifies an array of recipients of a message.
	Recipients *RecipientsArrayOfSmtpAddressType `xml:"m:Recipients"`
	// The SearchDumpster element specifies whether to search in the Exchange Dumpster.
	SearchDumpster *SearchDumpster `xml:"t:SearchDumpster"`
	// The Senders element specifies an array of Simple Mail Transfer Protocol (SMTP) addresses.
	Senders *Senders `xml:"t:Senders"`
	// The ToDate element specifies the date that the message was received.
	ToDate *ToDate `xml:"t:ToDate"`
}
