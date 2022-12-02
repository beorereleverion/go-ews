package elements

// The Type element specifies the type of folder used in a retention policy.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/type-elcfoldertype
type TypeElcFolderType string

const (
	// All
	TypeElcFolderTypeAll TypeElcFolderType = `All`
	// Calendar
	TypeElcFolderTypeCalendar TypeElcFolderType = `Calendar`
	// Contacts
	TypeElcFolderTypeContacts TypeElcFolderType = `Contacts`
	// ConversationHistory
	TypeElcFolderTypeConversationHistory TypeElcFolderType = `ConversationHistory`
	// DeletedItems
	TypeElcFolderTypeDeletedItems TypeElcFolderType = `DeletedItems`
	// Drafts
	TypeElcFolderTypeDrafts TypeElcFolderType = `Drafts`
	// Inbox
	TypeElcFolderTypeInbox TypeElcFolderType = `Inbox`
	// Journal
	TypeElcFolderTypeJournal TypeElcFolderType = `Journal`
	// JunkEmail
	TypeElcFolderTypeJunkEmail TypeElcFolderType = `JunkEmail`
	// ManagedCustomFolder
	TypeElcFolderTypeManagedCustomFolder TypeElcFolderType = `ManagedCustomFolder`
	// NonIpmRoot
	TypeElcFolderTypeNonIpmRoot TypeElcFolderType = `NonIpmRoot`
	// Notes
	TypeElcFolderTypeNotes TypeElcFolderType = `Notes`
	// Outbox
	TypeElcFolderTypeOutbox TypeElcFolderType = `Outbox`
	// Personal
	TypeElcFolderTypePersonal TypeElcFolderType = `Personal`
	// RecoverableItems
	TypeElcFolderTypeRecoverableItems TypeElcFolderType = `RecoverableItems`
	// RssSubscriptions
	TypeElcFolderTypeRssSubscriptions TypeElcFolderType = `RssSubscriptions`
	// SentItems
	TypeElcFolderTypeSentItems TypeElcFolderType = `SentItems`
	// SyncIssues
	TypeElcFolderTypeSyncIssues TypeElcFolderType = `SyncIssues`
	// Tasks
	TypeElcFolderTypeTasks TypeElcFolderType = `Tasks`
)
