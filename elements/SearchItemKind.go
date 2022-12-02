package elements

// The SearchItemKind element indicates the type of items that are searched for a FindMailboxStatisticsByKeyword operation.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/searchitemkind
type SearchItemKind string

const (
	// Contacts - Indicates that contacts are searched for keywords.
	SearchItemKindContacts SearchItemKind = `Contacts`
	// Docs - Indicates that documents are searched for keywords.
	SearchItemKindDocs SearchItemKind = `Docs`
	// Email - Indicates that email messages are searched for keywords.
	SearchItemKindEmail SearchItemKind = `Email`
	// Faxes - Indicates that faxes are searched for keywords.
	SearchItemKindFaxes SearchItemKind = `Faxes`
	// Im - Indicates that instant messages are searched for keywords.
	SearchItemKindIm SearchItemKind = `Im`
	// Journals - Indicates that journals are searched for keywords.
	SearchItemKindJournals SearchItemKind = `Journals`
	// Meetings - Indicates that meetings are searched for keywords.
	SearchItemKindMeetings SearchItemKind = `Meetings`
	// Notes - Indicates that notes are searched for keywords.
	SearchItemKindNotes SearchItemKind = `Notes`
	// Posts - Indicates that posts are searched for keywords.
	SearchItemKindPosts SearchItemKind = `Posts`
	// Rssfeeds - Indicates that RSS feeds are searched for keywords.
	SearchItemKindRssfeeds SearchItemKind = `Rssfeeds`
	// Tasks - Indicates that tasks are searched for keywords.
	SearchItemKindTasks SearchItemKind = `Tasks`
	// Voicemail - Indicates that voice mails are searched for keywords.
	SearchItemKindVoicemail SearchItemKind = `Voicemail`
)
