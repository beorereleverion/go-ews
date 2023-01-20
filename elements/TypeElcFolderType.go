package elements

// The Type element specifies the type of folder used in a retention policy.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/type-elcfoldertype
import "encoding/xml"

type TypeElcFolderType struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

const (
	// All
	TypeElcFolderTypeAll string = `All`
	// Calendar
	TypeElcFolderTypeCalendar string = `Calendar`
	// Contacts
	TypeElcFolderTypeContacts string = `Contacts`
	// ConversationHistory
	TypeElcFolderTypeConversationHistory string = `ConversationHistory`
	// DeletedItems
	TypeElcFolderTypeDeletedItems string = `DeletedItems`
	// Drafts
	TypeElcFolderTypeDrafts string = `Drafts`
	// Inbox
	TypeElcFolderTypeInbox string = `Inbox`
	// Journal
	TypeElcFolderTypeJournal string = `Journal`
	// JunkEmail
	TypeElcFolderTypeJunkEmail string = `JunkEmail`
	// ManagedCustomFolder
	TypeElcFolderTypeManagedCustomFolder string = `ManagedCustomFolder`
	// NonIpmRoot
	TypeElcFolderTypeNonIpmRoot string = `NonIpmRoot`
	// Notes
	TypeElcFolderTypeNotes string = `Notes`
	// Outbox
	TypeElcFolderTypeOutbox string = `Outbox`
	// Personal
	TypeElcFolderTypePersonal string = `Personal`
	// RecoverableItems
	TypeElcFolderTypeRecoverableItems string = `RecoverableItems`
	// RssSubscriptions
	TypeElcFolderTypeRssSubscriptions string = `RssSubscriptions`
	// SentItems
	TypeElcFolderTypeSentItems string = `SentItems`
	// SyncIssues
	TypeElcFolderTypeSyncIssues string = `SyncIssues`
	// Tasks
	TypeElcFolderTypeTasks string = `Tasks`
)

func (T *TypeElcFolderType) SetForMarshal() {
	T.XMLName.Local = "t:Type"
}

func (T *TypeElcFolderType) GetSchema() *Schema {
	return &SchemaTypes
}
