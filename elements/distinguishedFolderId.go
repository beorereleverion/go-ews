package elements

// The DistinguishedFolderId element identifies folders that can be referenced by name. If you do not use this element, you must use the FolderId element to identify a folder.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/distinguishedfolderid
type DistinguishedFolderId struct {
	// Contains a string that identifies a default folder. This attribute is required.
	ID DistinguishedFolderIdAttrID `xml:"Id,attr"`
	// Contains a string that identifies a version of a folder that is identified by the Id attribute. This attribute is optional. Use this attribute to make sure that the correct version of a folder is used.
	ChangeKey *string `xml:"ChangeKey,attr"`
	// Identifies a primary SMTP address. Proxy addresses are not allowed.
	Mailbox *string `xml:"t:Mailbox"`
}

type DistinguishedFolderIdAttrID string

const (
	// Represents the Calendar folder.
	DistinguishedFolderIdAttrIDcalendar DistinguishedFolderIdAttrID = `calendar`
	// Represents the Contacts folder.
	DistinguishedFolderIdAttrIDcontacts DistinguishedFolderIdAttrID = `contacts`
	// Represents the Deleted Items folder.
	DistinguishedFolderIdAttrIDdeleteditems DistinguishedFolderIdAttrID = `deleteditems`
	// Represents the Drafts folder.
	DistinguishedFolderIdAttrIDdrafts DistinguishedFolderIdAttrID = `drafts`
	// Represents the Inbox folder.
	DistinguishedFolderIdAttrIDinbox DistinguishedFolderIdAttrID = `inbox`
	// Represents the Journal folder.
	DistinguishedFolderIdAttrIDjournal DistinguishedFolderIdAttrID = `journal`
	// Represents the Notes folder.
	DistinguishedFolderIdAttrIDnotes DistinguishedFolderIdAttrID = `notes`
	// Represents the Outbox folder.
	DistinguishedFolderIdAttrIDoutbox DistinguishedFolderIdAttrID = `outbox`
	// Represents the Sent Items folder.
	DistinguishedFolderIdAttrIDsentitems DistinguishedFolderIdAttrID = `sentitems`
	// Represents the Tasks folder.
	DistinguishedFolderIdAttrIDtasks DistinguishedFolderIdAttrID = `tasks`
	// Represents the message folder root.
	DistinguishedFolderIdAttrIDmsgfolderroot DistinguishedFolderIdAttrID = `msgfolderroot`
	// Represents the root of the mailbox.
	DistinguishedFolderIdAttrIDroot DistinguishedFolderIdAttrID = `root`
	// Represents the Junk Email folder.
	DistinguishedFolderIdAttrIDjunkemail DistinguishedFolderIdAttrID = `junkemail`
	// Represents the Search Folders folder.
	DistinguishedFolderIdAttrIDsearchfolders DistinguishedFolderIdAttrID = `searchfolders`
	// Represents the Voice Mail folder.
	DistinguishedFolderIdAttrIDvoicemail DistinguishedFolderIdAttrID = `voicemail`
	// Represents the dumpster root folder.
	DistinguishedFolderIdAttrIDrecoverableitemsroot DistinguishedFolderIdAttrID = `recoverableitemsroot`
	// Represents the dumpster deletions folder.
	DistinguishedFolderIdAttrIDrecoverableitemsdeletions DistinguishedFolderIdAttrID = `recoverableitemsdeletions`
	// Represents the dumpster versions folder.
	DistinguishedFolderIdAttrIDrecoverableitemsversions DistinguishedFolderIdAttrID = `recoverableitemsversions`
	// Represents the dumpster purges folder.
	DistinguishedFolderIdAttrIDrecoverableitemspurges DistinguishedFolderIdAttrID = `recoverableitemspurges`
	// Represents the root archive folder.
	DistinguishedFolderIdAttrIDarchiveroot DistinguishedFolderIdAttrID = `archiveroot`
	// Represents the root archive message folder.
	DistinguishedFolderIdAttrIDarchivemsgfolderroot DistinguishedFolderIdAttrID = `archivemsgfolderroot`
	// Represents the archive deleted items folder.
	DistinguishedFolderIdAttrIDarchivedeleteditems DistinguishedFolderIdAttrID = `archivedeleteditems`
	// Represents the archive Inbox folder. Versions of Exchange starting with build number 15.00.0913.09 include this value.
	DistinguishedFolderIdAttrIDarchiveinbox DistinguishedFolderIdAttrID = `archiveinbox`
	// Represents the archive recoverable items root folder.
	DistinguishedFolderIdAttrIDarchiverecoverableitemsroot DistinguishedFolderIdAttrID = `archiverecoverableitemsroot`
	// Represents the archive recoverable items deletions folder.
	DistinguishedFolderIdAttrIDarchiverecoverableitemsdeletions DistinguishedFolderIdAttrID = `archiverecoverableitemsdeletions`
	// Represents the archive recoverable items versions folder.
	DistinguishedFolderIdAttrIDarchiverecoverableitemsversions DistinguishedFolderIdAttrID = `archiverecoverableitemsversions`
	// Represents the archive recoverable items purges folder.
	DistinguishedFolderIdAttrIDarchiverecoverableitemspurges DistinguishedFolderIdAttrID = `archiverecoverableitemspurges`
	// Represents the sync issues folder.
	DistinguishedFolderIdAttrIDsyncissues DistinguishedFolderIdAttrID = `syncissues`
	// Represents the conflicts folder.
	DistinguishedFolderIdAttrIDconflicts DistinguishedFolderIdAttrID = `conflicts`
	// Represents the local failures folder.
	DistinguishedFolderIdAttrIDlocalfailures DistinguishedFolderIdAttrID = `localfailures`
	// Represents the server failures folder.
	DistinguishedFolderIdAttrIDserverfailures DistinguishedFolderIdAttrID = `serverfailures`
	// Represents the recipient cache folder.
	DistinguishedFolderIdAttrIDrecipientcache DistinguishedFolderIdAttrID = `recipientcache`
	// Represents the quick contacts folder.
	DistinguishedFolderIdAttrIDquickcontacts DistinguishedFolderIdAttrID = `quickcontacts`
	// Represents the conversation history folder.
	DistinguishedFolderIdAttrIDconversationhistory DistinguishedFolderIdAttrID = `conversationhistory`
	// Represents the admin audit logs folder.
	DistinguishedFolderIdAttrIDadminauditlogs DistinguishedFolderIdAttrID = `adminauditlogs`
	// Represents the todo search folder.
	DistinguishedFolderIdAttrIDtodosearch DistinguishedFolderIdAttrID = `todosearch`
	// Represents the My Contacts folder.
	DistinguishedFolderIdAttrIDmycontacts DistinguishedFolderIdAttrID = `mycontacts`
	// Represents the directory folder.
	DistinguishedFolderIdAttrIDdirectory DistinguishedFolderIdAttrID = `directory`
	// Represents the IM contact list folder.
	DistinguishedFolderIdAttrIDimcontactlist DistinguishedFolderIdAttrID = `imcontactlist`
	// Represents the people connect folder.
	DistinguishedFolderIdAttrIDpeopleconnect DistinguishedFolderIdAttrID = `peopleconnect`
	// Represents the Favorites folder.
	DistinguishedFolderIdAttrIDfavorites DistinguishedFolderIdAttrID = `favorites`
)
