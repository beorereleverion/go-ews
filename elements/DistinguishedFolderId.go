package elements

// The DistinguishedFolderId element identifies folders that can be referenced by name. If you do not use this element, you must use the FolderId element to identify a folder.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/distinguishedfolderid
type DistinguishedFolderId struct {
	// The Mailbox element identifies a mail-enabled Active Directory object.
	Mailbox *Mailbox `xml:"t:Mailbox"`
	// Contains a string that identifies a version of a folder that is identified by the Id attribute. This attribute is optional. Use this attribute to make sure that the correct version of a folder is used.
	ChangeKey *string `xml:"ChangeKey,attr"`
	// Contains a string that identifies a default folder. This attribute is required.
	Id *string `xml:"Id,attr"`
}

const (
	// Represents the Calendar folder.
	DistinguishedFolderIdcalendar = `calendar`
	// Represents the Contacts folder.
	DistinguishedFolderIdcontacts = `contacts`
	// Represents the Deleted Items folder.
	DistinguishedFolderIddeleteditems = `deleteditems`
	// Represents the Drafts folder.
	DistinguishedFolderIddrafts = `drafts`
	// Represents the Inbox folder.
	DistinguishedFolderIdinbox = `inbox`
	// Represents the Journal folder.
	DistinguishedFolderIdjournal = `journal`
	// Represents the Notes folder.
	DistinguishedFolderIdnotes = `notes`
	// Represents the Outbox folder.
	DistinguishedFolderIdoutbox = `outbox`
	// Represents the Sent Items folder.
	DistinguishedFolderIdsentitems = `sentitems`
	// Represents the Tasks folder.
	DistinguishedFolderIdtasks = `tasks`
	// Represents the message folder root.
	DistinguishedFolderIdmsgfolderroot = `msgfolderroot`
	// Represents the root of the mailbox.
	DistinguishedFolderIdroot = `root`
	// Represents the Junk Email folder.
	DistinguishedFolderIdjunkemail = `junkemail`
	// Represents the Search Folders folder.
	DistinguishedFolderIdsearchfolders = `searchfolders`
	// Represents the Voice Mail folder.
	DistinguishedFolderIdvoicemail = `voicemail`
	// Represents the dumpster root folder.
	DistinguishedFolderIdrecoverableitemsroot = `recoverableitemsroot`
	// Represents the dumpster deletions folder.
	DistinguishedFolderIdrecoverableitemsdeletions = `recoverableitemsdeletions`
	// Represents the dumpster versions folder.
	DistinguishedFolderIdrecoverableitemsversions = `recoverableitemsversions`
	// Represents the dumpster purges folder.
	DistinguishedFolderIdrecoverableitemspurges = `recoverableitemspurges`
	// Represents the root archive folder.
	DistinguishedFolderIdarchiveroot = `archiveroot`
	// Represents the root archive message folder.
	DistinguishedFolderIdarchivemsgfolderroot = `archivemsgfolderroot`
	// Represents the archive deleted items folder.
	DistinguishedFolderIdarchivedeleteditems = `archivedeleteditems`
	// Represents the archive Inbox folder. Versions of Exchange starting with build number 15.00.0913.09 include this value.
	DistinguishedFolderIdarchiveinbox = `archiveinbox`
	// Represents the archive recoverable items root folder.
	DistinguishedFolderIdarchiverecoverableitemsroot = `archiverecoverableitemsroot`
	// Represents the archive recoverable items deletions folder.
	DistinguishedFolderIdarchiverecoverableitemsdeletions = `archiverecoverableitemsdeletions`
	// Represents the archive recoverable items versions folder.
	DistinguishedFolderIdarchiverecoverableitemsversions = `archiverecoverableitemsversions`
	// Represents the archive recoverable items purges folder.
	DistinguishedFolderIdarchiverecoverableitemspurges = `archiverecoverableitemspurges`
	// Represents the sync issues folder.
	DistinguishedFolderIdsyncissues = `syncissues`
	// Represents the conflicts folder.
	DistinguishedFolderIdconflicts = `conflicts`
	// Represents the local failures folder.
	DistinguishedFolderIdlocalfailures = `localfailures`
	// Represents the server failures folder.
	DistinguishedFolderIdserverfailures = `serverfailures`
	// Represents the recipient cache folder.
	DistinguishedFolderIdrecipientcache = `recipientcache`
	// Represents the quick contacts folder.
	DistinguishedFolderIdquickcontacts = `quickcontacts`
	// Represents the conversation history folder.
	DistinguishedFolderIdconversationhistory = `conversationhistory`
	// Represents the admin audit logs folder.
	DistinguishedFolderIdadminauditlogs = `adminauditlogs`
	// Represents the todo search folder.
	DistinguishedFolderIdtodosearch = `todosearch`
	// Represents the My Contacts folder.
	DistinguishedFolderIdmycontacts = `mycontacts`
	// Represents the directory folder.
	DistinguishedFolderIddirectory = `directory`
	// Represents the IM contact list folder.
	DistinguishedFolderIdimcontactlist = `imcontactlist`
	// Represents the people connect folder.
	DistinguishedFolderIdpeopleconnect = `peopleconnect`
	// Represents the Favorites folder.
	DistinguishedFolderIdfavorites = `favorites`
)
