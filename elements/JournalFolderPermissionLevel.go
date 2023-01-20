package elements

// The JournalFolderPermissionLevel element contains the permissions for the default Journal folder. This element was introduced in Microsoft Exchange Server 2007 Service Pack 1 (SP1).
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/journalfolderpermissionlevel
import "encoding/xml"

type JournalFolderPermissionLevel struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

const (
	// The delegate user can read and create items in the Journal folder.
	JournalFolderPermissionLevelAuthor string = `Author`
	// The delegate user has custom access permissions to the Journal folder.
	JournalFolderPermissionLevelCustom string = `Custom`
	// The delegate user can read, create, and modify items in the Journal folder.
	JournalFolderPermissionLevelEditor string = `Editor`
	// The delegate user has no access permissions to the Journal folder.
	JournalFolderPermissionLevelNone string = `None`
	// The delegate user can read items in the Journal folder.
	JournalFolderPermissionLevelReviewer string = `Reviewer`
)

func (J *JournalFolderPermissionLevel) SetForMarshal() {
	J.XMLName.Local = "t:JournalFolderPermissionLevel"
}

func (J *JournalFolderPermissionLevel) GetSchema() *Schema {
	return &SchemaTypes
}
