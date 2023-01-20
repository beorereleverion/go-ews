package elements

// The TasksFolderPermissionLevel element contains the permissions for the default Tasks folder. This element was introduced in Microsoft Exchange Server 2007 Service Pack 1 (SP1).
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/tasksfolderpermissionlevel
import "encoding/xml"

type TasksFolderPermissionLevel struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

const (
	// The delegate user can read and create items in the Tasks folder.
	TasksFolderPermissionLevelAuthor string = `Author`
	// The delegate user has custom access permissions to the Tasks folder.
	TasksFolderPermissionLevelCustom string = `Custom`
	// The delegate user can read, create, and modify items in the Tasks folder.
	TasksFolderPermissionLevelEditor string = `Editor`
	// The delegate user has no access permissions to the Tasks folder.
	TasksFolderPermissionLevelNone string = `None`
	// The delegate user can read items in the Tasks folder.
	TasksFolderPermissionLevelReviewer string = `Reviewer`
)

func (T *TasksFolderPermissionLevel) SetForMarshal() {
	T.XMLName.Local = "t:TasksFolderPermissionLevel"
}

func (T *TasksFolderPermissionLevel) GetSchema() *Schema {
	return &SchemaTypes
}
