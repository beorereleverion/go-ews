package elements

// The NotesFolderPermissionLevel element contains the permissions for the default Notes folder. This element was introduced in Microsoft Exchange Server 2007 Service Pack 1 (SP1).
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/notesfolderpermissionlevel
import "encoding/xml"

type NotesFolderPermissionLevel struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

const (
	// The delegate user can read and create items in the Notes folder.
	NotesFolderPermissionLevelAuthor string = `Author`
	// The delegate user has custom access permissions to the Notes folder.
	NotesFolderPermissionLevelCustom string = `Custom`
	// The delegate user can read, create, and modify items in the Notes folder.
	NotesFolderPermissionLevelEditor string = `Editor`
	// The delegate user has no access permissions to the Notes folder.
	NotesFolderPermissionLevelNone string = `None`
	// The delegate user can read items in the Notes folder.
	NotesFolderPermissionLevelReviewer string = `Reviewer`
)

func (N *NotesFolderPermissionLevel) SetForMarshal() {
	N.XMLName.Local = "t:NotesFolderPermissionLevel"
}

func (N *NotesFolderPermissionLevel) GetSchema() *Schema {
	return &SchemaTypes
}
