package elements

// The ContactsFolderPermissionLevel element contains the permissions for the default Contacts folder. This element was introduced in Microsoft Exchange Server 2007 Service Pack 1 (SP1).
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/contactsfolderpermissionlevel
import "encoding/xml"

type ContactsFolderPermissionLevel struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

const (
	// The delegate user can read and create items in the Contacts folder.
	ContactsFolderPermissionLevelAuthor string = `Author`
	// The delegate user has custom access permissions to the Contacts folder.
	ContactsFolderPermissionLevelCustom string = `Custom`
	// The delegate user can read, create, and modify items in the Contacts folder.
	ContactsFolderPermissionLevelEditor string = `Editor`
	// The delegate user has no access permissions to the Contacts folder.
	ContactsFolderPermissionLevelNone string = `None`
	// The delegate user can read items in the Contacts folder.
	ContactsFolderPermissionLevelReviewer string = `Reviewer`
)

func (C *ContactsFolderPermissionLevel) SetForMarshal() {
	C.XMLName.Local = "t:ContactsFolderPermissionLevel"
}

func (C *ContactsFolderPermissionLevel) GetSchema() *Schema {
	return &SchemaTypes
}
