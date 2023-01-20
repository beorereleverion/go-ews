package elements

// The PermissionLevel element represents the permission level that a user has on a folder. This element was introduced in Microsoft Exchange Server 2007 Service Pack 1 (SP1).
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/permissionlevel
import "encoding/xml"

type PermissionLevel struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

const (
	// Author
	PermissionLevelAuthor string = `Author`
	// Contributor
	PermissionLevelContributor string = `Contributor`
	// Custom
	PermissionLevelCustom string = `Custom`
	// Editor
	PermissionLevelEditor string = `Editor`
	// None
	PermissionLevelNone string = `None`
	// NoneditingAuthor
	PermissionLevelNoneditingAuthor string = `NoneditingAuthor`
	// Owner
	PermissionLevelOwner string = `Owner`
	// PublishingAuthor
	PermissionLevelPublishingAuthor string = `PublishingAuthor`
	// PublishingEditor
	PermissionLevelPublishingEditor string = `PublishingEditor`
	// Reviewer
	PermissionLevelReviewer string = `Reviewer`
)

func (P *PermissionLevel) SetForMarshal() {
	P.XMLName.Local = "t:PermissionLevel"
}

func (P *PermissionLevel) GetSchema() *Schema {
	return &SchemaTypes
}
