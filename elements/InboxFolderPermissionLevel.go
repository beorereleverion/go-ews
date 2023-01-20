package elements

// The InboxFolderPermissionLevel element contains the permissions for the default Inbox folder. This element was introduced in Microsoft Exchange Server 2007 Service Pack 1 (SP1).
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/inboxfolderpermissionlevel
import "encoding/xml"

type InboxFolderPermissionLevel struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

const (
	// The delegate user can read and create items in the Inbox folder.
	InboxFolderPermissionLevelAuthor string = `Author`
	// The delegate user has custom access permissions to the Inbox folder.
	InboxFolderPermissionLevelCustom string = `Custom`
	// The delegate user can read, create, and modify items in the Inbox folder.
	InboxFolderPermissionLevelEditor string = `Editor`
	// The delegate user has no access permissions to the Inbox folder.
	InboxFolderPermissionLevelNone string = `None`
	// The delegate user can read items in the Inbox folder.
	InboxFolderPermissionLevelReviewer string = `Reviewer`
)

func (I *InboxFolderPermissionLevel) SetForMarshal() {
	I.XMLName.Local = "t:InboxFolderPermissionLevel"
}

func (I *InboxFolderPermissionLevel) GetSchema() *Schema {
	return &SchemaTypes
}
