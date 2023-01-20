package elements

// The SearchItemKind element indicates the type of items that are searched for a FindMailboxStatisticsByKeyword operation.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/searchitemkind
import "encoding/xml"

type SearchItemKind struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

const (
	// Contacts - Indicates that contacts are searched for keywords.
	SearchItemKindContacts string = `Contacts`
	// Docs - Indicates that documents are searched for keywords.
	SearchItemKindDocs string = `Docs`
	// Email - Indicates that email messages are searched for keywords.
	SearchItemKindEmail string = `Email`
	// Faxes - Indicates that faxes are searched for keywords.
	SearchItemKindFaxes string = `Faxes`
	// Im - Indicates that instant messages are searched for keywords.
	SearchItemKindIm string = `Im`
	// Journals - Indicates that journals are searched for keywords.
	SearchItemKindJournals string = `Journals`
	// Meetings - Indicates that meetings are searched for keywords.
	SearchItemKindMeetings string = `Meetings`
	// Notes - Indicates that notes are searched for keywords.
	SearchItemKindNotes string = `Notes`
	// Posts - Indicates that posts are searched for keywords.
	SearchItemKindPosts string = `Posts`
	// Rssfeeds - Indicates that RSS feeds are searched for keywords.
	SearchItemKindRssfeeds string = `Rssfeeds`
	// Tasks - Indicates that tasks are searched for keywords.
	SearchItemKindTasks string = `Tasks`
	// Voicemail - Indicates that voice mails are searched for keywords.
	SearchItemKindVoicemail string = `Voicemail`
)

func (S *SearchItemKind) SetForMarshal() {
	S.XMLName.Local = "t:SearchItemKind"
}

func (S *SearchItemKind) GetSchema() *Schema {
	return &SchemaTypes
}
