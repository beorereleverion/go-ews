package elements

// The DistinguishedFolderId element identifies folders that can be referenced by name.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/distinguishedfolderid-distinguishedfolderidnametype
type DistinguishedFolderIdDistinguishedFolderIdNameType string

const (
	// Indicates the URL of the administrative audit log folder.
	DistinguishedFolderIdDistinguishedFolderIdNameTypeadminauditlogs DistinguishedFolderIdDistinguishedFolderIdNameType = `adminauditlogs`
	// Indicates the URL of the archived deleted items folder.
	DistinguishedFolderIdDistinguishedFolderIdNameTypearchivedeleteditems DistinguishedFolderIdDistinguishedFolderIdNameType = `archivedeleteditems`
	// Indicates the URL of the archived message folder root folder.
	DistinguishedFolderIdDistinguishedFolderIdNameTypearchivemsgfolderroot DistinguishedFolderIdDistinguishedFolderIdNameType = `archivemsgfolderroot`
	// Indicates the URL of the archived recoverable deleted items folder.
	DistinguishedFolderIdDistinguishedFolderIdNameTypearchiverecoverableitemsdeletions DistinguishedFolderIdDistinguishedFolderIdNameType = `archiverecoverableitemsdeletions`
	// Indicates the URL of the archived purged recoverable items folder.
	DistinguishedFolderIdDistinguishedFolderIdNameTypearchiverecoverableitemspurges DistinguishedFolderIdDistinguishedFolderIdNameType = `archiverecoverableitemspurges`
	// Indicates the URL of the archived recoverable items root folder.
	DistinguishedFolderIdDistinguishedFolderIdNameTypearchiverecoverableitemsroot DistinguishedFolderIdDistinguishedFolderIdNameType = `archiverecoverableitemsroot`
	// Indicates the URL of the archived recoverable items versions folder.
	DistinguishedFolderIdDistinguishedFolderIdNameTypearchiverecoverableitemsversions DistinguishedFolderIdDistinguishedFolderIdNameType = `archiverecoverableitemsversions`
	// Indicates the URL of the archive root folder.
	DistinguishedFolderIdDistinguishedFolderIdNameTypearchiveroot DistinguishedFolderIdDistinguishedFolderIdNameType = `archiveroot`
	// Indicates the URL of the calendar folder.
	DistinguishedFolderIdDistinguishedFolderIdNameTypecalendar DistinguishedFolderIdDistinguishedFolderIdNameType = `calendar`
	// Indicates the URL of the conflicts folder.
	DistinguishedFolderIdDistinguishedFolderIdNameTypeconflicts DistinguishedFolderIdDistinguishedFolderIdNameType = `conflicts`
	// Indicates the URL of the contacts folder.
	DistinguishedFolderIdDistinguishedFolderIdNameTypecontacts DistinguishedFolderIdDistinguishedFolderIdNameType = `contacts`
	// Indicates the URL of the conversation history folder.
	DistinguishedFolderIdDistinguishedFolderIdNameTypeconversationhistory DistinguishedFolderIdDistinguishedFolderIdNameType = `conversationhistory`
	// Indicates the URL of the deleted items folder.
	DistinguishedFolderIdDistinguishedFolderIdNameTypedeleteditems DistinguishedFolderIdDistinguishedFolderIdNameType = `deleteditems`
	// Indicates a URL of the directory folder.
	DistinguishedFolderIdDistinguishedFolderIdNameTypedirectory DistinguishedFolderIdDistinguishedFolderIdNameType = `directory`
	// Indicates the URL of the drafts folder.
	DistinguishedFolderIdDistinguishedFolderIdNameTypedrafts DistinguishedFolderIdDistinguishedFolderIdNameType = `drafts`
	// Indicates the URL of the inbox folder.
	DistinguishedFolderIdDistinguishedFolderIdNameTypeinbox DistinguishedFolderIdDistinguishedFolderIdNameType = `inbox`
	// Indicates the URL of the journal folder.
	DistinguishedFolderIdDistinguishedFolderIdNameTypejournal DistinguishedFolderIdDistinguishedFolderIdNameType = `journal`
	// Indicates the URL of the junk email folder.
	DistinguishedFolderIdDistinguishedFolderIdNameTypejunkemail DistinguishedFolderIdDistinguishedFolderIdNameType = `junkemail`
	// Indicates the URL of the local failures folder.
	DistinguishedFolderIdDistinguishedFolderIdNameTypelocalfailures DistinguishedFolderIdDistinguishedFolderIdNameType = `localfailures`
	// Indicates the URL of the message root folder.
	DistinguishedFolderIdDistinguishedFolderIdNameTypemsgfolderroot DistinguishedFolderIdDistinguishedFolderIdNameType = `msgfolderroot`
	// Indicates the URL of the my contacts folder.
	DistinguishedFolderIdDistinguishedFolderIdNameTypemycontacts DistinguishedFolderIdDistinguishedFolderIdNameType = `mycontacts`
	// Indicates the URL of the notes folder.
	DistinguishedFolderIdDistinguishedFolderIdNameTypenotes DistinguishedFolderIdDistinguishedFolderIdNameType = `notes`
	// Indicates the URL of the outbox folder.
	DistinguishedFolderIdDistinguishedFolderIdNameTypeoutbox DistinguishedFolderIdDistinguishedFolderIdNameType = `outbox`
	// Indicates the URL of the public folders root folder.
	DistinguishedFolderIdDistinguishedFolderIdNameTypepublicfoldersroot DistinguishedFolderIdDistinguishedFolderIdNameType = `publicfoldersroot`
	// Indicates the URL of the quick contacts folder.
	DistinguishedFolderIdDistinguishedFolderIdNameTypequickcontacts DistinguishedFolderIdDistinguishedFolderIdNameType = `quickcontacts`
	// Indicates the URL of the recipient cache folder.
	DistinguishedFolderIdDistinguishedFolderIdNameTyperecipientcache DistinguishedFolderIdDistinguishedFolderIdNameType = `recipientcache`
	// Indicates the URL of the deleted recoverable items folder.
	DistinguishedFolderIdDistinguishedFolderIdNameTyperecoverableitemsdeletions DistinguishedFolderIdDistinguishedFolderIdNameType = `recoverableitemsdeletions`
	// Indicates the URL of the purged recoverable items folder.
	DistinguishedFolderIdDistinguishedFolderIdNameTyperecoverableitemspurges DistinguishedFolderIdDistinguishedFolderIdNameType = `recoverableitemspurges`
	// Indicates the URL of the recoverable items root folder.
	DistinguishedFolderIdDistinguishedFolderIdNameTyperecoverableitemsroot DistinguishedFolderIdDistinguishedFolderIdNameType = `recoverableitemsroot`
	// Indicates the URL of the recoverable item versions folder.
	DistinguishedFolderIdDistinguishedFolderIdNameTyperecoverableitemsversions DistinguishedFolderIdDistinguishedFolderIdNameType = `recoverableitemsversions`
	// Indicates the URL of the root folder.
	DistinguishedFolderIdDistinguishedFolderIdNameTyperoot DistinguishedFolderIdDistinguishedFolderIdNameType = `root`
	// Indicates the URL of the search folders.
	DistinguishedFolderIdDistinguishedFolderIdNameTypesearchfolders DistinguishedFolderIdDistinguishedFolderIdNameType = `searchfolders`
	// Indicates the URL of the sent items folder.
	DistinguishedFolderIdDistinguishedFolderIdNameTypesentitems DistinguishedFolderIdDistinguishedFolderIdNameType = `sentitems`
	// Indicates the URL of the server failures folder.
	DistinguishedFolderIdDistinguishedFolderIdNameTypeserverfailures DistinguishedFolderIdDistinguishedFolderIdNameType = `serverfailures`
	// Indicates the URL of the synchronization issues folder.
	DistinguishedFolderIdDistinguishedFolderIdNameTypesyncissues DistinguishedFolderIdDistinguishedFolderIdNameType = `syncissues`
	// Indicates the URL of the tasks folder.
	DistinguishedFolderIdDistinguishedFolderIdNameTypetasks DistinguishedFolderIdDistinguishedFolderIdNameType = `tasks`
	// Indicates the URL of the search to-do folder.
	DistinguishedFolderIdDistinguishedFolderIdNameTypetodosearch DistinguishedFolderIdDistinguishedFolderIdNameType = `todosearch`
	// Indicates the URL of the voice-mail folder.
	DistinguishedFolderIdDistinguishedFolderIdNameTypevoicemail DistinguishedFolderIdDistinguishedFolderIdNameType = `voicemail`
)
