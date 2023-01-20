package elements

// The FindMailboxStatisticsByKeywords element specifies a request to search for mailbox statistics by keyword.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/findmailboxstatisticsbykeywords
import "encoding/xml"

type FindMailboxStatisticsByKeywords struct {
	XMLName xml.Name

	// The FromDate element specifies the date that the message was sent.
	FromDate *FromDate `xml:"FromDate"`
	// The IncludePersonalArchive element specifies whether to include the personal archive in the search.
	IncludePersonalArchive *IncludePersonalArchive `xml:"IncludePersonalArchive"`
	// The IncludeUnsearchableItems element specifies whether to include items that cannot be searched.
	IncludeUnsearchableItems *IncludeUnsearchableItems `xml:"IncludeUnsearchableItems"`
	// The Keywords element specifies keywords for a FindMailboxStatisticsByKeywords operation search.
	Keywords *Keywords `xml:"Keywords"`
	// The Language element contains the language used for the search query.
	Language *Language `xml:"Language"`
	// The Mailboxes element contains an array of mailboxes.
	Mailboxes *MailboxesArrayOfUserMailboxesType `xml:"Mailboxes"`
	// The MessageTypes element specifies an array of messages to search.
	MessageTypes *MessageTypes `xml:"MessageTypes"`
	// The Recipients element specifies an array of recipients of a message.
	Recipients *RecipientsArrayOfSmtpAddressType `xml:"Recipients"`
	// The SearchDumpster element specifies whether to search in the Exchange Dumpster.
	SearchDumpster *SearchDumpster `xml:"SearchDumpster"`
	// The Senders element specifies an array of Simple Mail Transfer Protocol (SMTP) addresses.
	Senders *Senders `xml:"Senders"`
	// The ToDate element specifies the date that the message was received.
	ToDate *ToDate `xml:"ToDate"`
}

func (F *FindMailboxStatisticsByKeywords) SetForMarshal() {
	F.XMLName.Local = "m:FindMailboxStatisticsByKeywords"
}

func (F *FindMailboxStatisticsByKeywords) GetSchema() *Schema {
	return &SchemaMessages
}
