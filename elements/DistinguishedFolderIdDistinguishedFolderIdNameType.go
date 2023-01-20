package elements

// The DistinguishedFolderId element identifies folders that can be referenced by name.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/distinguishedfolderid-distinguishedfolderidnametype
import "encoding/xml"

type DistinguishedFolderIdDistinguishedFolderIdNameType struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

const (
	// Indicates the URL of the administrative audit log folder.
	DistinguishedFolderIdDistinguishedFolderIdNameTypeadminauditlogs string = `adminauditlogs`
	// Indicates the URL of the archived deleted items folder.
	DistinguishedFolderIdDistinguishedFolderIdNameTypearchivedeleteditems string = `archivedeleteditems`
	// Indicates the URL of the archived message folder root folder.
	DistinguishedFolderIdDistinguishedFolderIdNameTypearchivemsgfolderroot string = `archivemsgfolderroot`
	// Indicates the URL of the archived recoverable deleted items folder.
	DistinguishedFolderIdDistinguishedFolderIdNameTypearchiverecoverableitemsdeletions string = `archiverecoverableitemsdeletions`
	// Indicates the URL of the archived purged recoverable items folder.
	DistinguishedFolderIdDistinguishedFolderIdNameTypearchiverecoverableitemspurges string = `archiverecoverableitemspurges`
	// Indicates the URL of the archived recoverable items root folder.
	DistinguishedFolderIdDistinguishedFolderIdNameTypearchiverecoverableitemsroot string = `archiverecoverableitemsroot`
	// Indicates the URL of the archived recoverable items versions folder.
	DistinguishedFolderIdDistinguishedFolderIdNameTypearchiverecoverableitemsversions string = `archiverecoverableitemsversions`
	// Indicates the URL of the archive root folder.
	DistinguishedFolderIdDistinguishedFolderIdNameTypearchiveroot string = `archiveroot`
	// Indicates the URL of the calendar folder.
	DistinguishedFolderIdDistinguishedFolderIdNameTypecalendar string = `calendar`
	// Indicates the URL of the conflicts folder.
	DistinguishedFolderIdDistinguishedFolderIdNameTypeconflicts string = `conflicts`
	// Indicates the URL of the contacts folder.
	DistinguishedFolderIdDistinguishedFolderIdNameTypecontacts string = `contacts`
	// Indicates the URL of the conversation history folder.
	DistinguishedFolderIdDistinguishedFolderIdNameTypeconversationhistory string = `conversationhistory`
	// Indicates the URL of the deleted items folder.
	DistinguishedFolderIdDistinguishedFolderIdNameTypedeleteditems string = `deleteditems`
	// Indicates a URL of the directory folder.
	DistinguishedFolderIdDistinguishedFolderIdNameTypedirectory string = `directory`
	// Indicates the URL of the drafts folder.
	DistinguishedFolderIdDistinguishedFolderIdNameTypedrafts string = `drafts`
	// Indicates the URL of the inbox folder.
	DistinguishedFolderIdDistinguishedFolderIdNameTypeinbox string = `inbox`
	// Indicates the URL of the journal folder.
	DistinguishedFolderIdDistinguishedFolderIdNameTypejournal string = `journal`
	// Indicates the URL of the junk email folder.
	DistinguishedFolderIdDistinguishedFolderIdNameTypejunkemail string = `junkemail`
	// Indicates the URL of the local failures folder.
	DistinguishedFolderIdDistinguishedFolderIdNameTypelocalfailures string = `localfailures`
	// Indicates the URL of the message root folder.
	DistinguishedFolderIdDistinguishedFolderIdNameTypemsgfolderroot string = `msgfolderroot`
	// Indicates the URL of the my contacts folder.
	DistinguishedFolderIdDistinguishedFolderIdNameTypemycontacts string = `mycontacts`
	// Indicates the URL of the notes folder.
	DistinguishedFolderIdDistinguishedFolderIdNameTypenotes string = `notes`
	// Indicates the URL of the outbox folder.
	DistinguishedFolderIdDistinguishedFolderIdNameTypeoutbox string = `outbox`
	// Indicates the URL of the public folders root folder.
	DistinguishedFolderIdDistinguishedFolderIdNameTypepublicfoldersroot string = `publicfoldersroot`
	// Indicates the URL of the quick contacts folder.
	DistinguishedFolderIdDistinguishedFolderIdNameTypequickcontacts string = `quickcontacts`
	// Indicates the URL of the recipient cache folder.
	DistinguishedFolderIdDistinguishedFolderIdNameTyperecipientcache string = `recipientcache`
	// Indicates the URL of the deleted recoverable items folder.
	DistinguishedFolderIdDistinguishedFolderIdNameTyperecoverableitemsdeletions string = `recoverableitemsdeletions`
	// Indicates the URL of the purged recoverable items folder.
	DistinguishedFolderIdDistinguishedFolderIdNameTyperecoverableitemspurges string = `recoverableitemspurges`
	// Indicates the URL of the recoverable items root folder.
	DistinguishedFolderIdDistinguishedFolderIdNameTyperecoverableitemsroot string = `recoverableitemsroot`
	// Indicates the URL of the recoverable item versions folder.
	DistinguishedFolderIdDistinguishedFolderIdNameTyperecoverableitemsversions string = `recoverableitemsversions`
	// Indicates the URL of the root folder.
	DistinguishedFolderIdDistinguishedFolderIdNameTyperoot string = `root`
	// Indicates the URL of the search folders.
	DistinguishedFolderIdDistinguishedFolderIdNameTypesearchfolders string = `searchfolders`
	// Indicates the URL of the sent items folder.
	DistinguishedFolderIdDistinguishedFolderIdNameTypesentitems string = `sentitems`
	// Indicates the URL of the server failures folder.
	DistinguishedFolderIdDistinguishedFolderIdNameTypeserverfailures string = `serverfailures`
	// Indicates the URL of the synchronization issues folder.
	DistinguishedFolderIdDistinguishedFolderIdNameTypesyncissues string = `syncissues`
	// Indicates the URL of the tasks folder.
	DistinguishedFolderIdDistinguishedFolderIdNameTypetasks string = `tasks`
	// Indicates the URL of the search to-do folder.
	DistinguishedFolderIdDistinguishedFolderIdNameTypetodosearch string = `todosearch`
	// Indicates the URL of the voice-mail folder.
	DistinguishedFolderIdDistinguishedFolderIdNameTypevoicemail string = `voicemail`
)

func (D *DistinguishedFolderIdDistinguishedFolderIdNameType) SetForMarshal() {
	D.XMLName.Local = "t:DistinguishedFolderId"
}

func (D *DistinguishedFolderIdDistinguishedFolderIdNameType) GetSchema() *Schema {
	return &SchemaTypes
}
