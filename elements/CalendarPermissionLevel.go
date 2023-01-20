package elements

// The CalendarPermissionLevel element represents the permission level that a user has on a Calendar folder. This element was introduced in Microsoft Exchange Server 2007 Service Pack 1 (SP1).
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/calendarpermissionlevel
import "encoding/xml"

type CalendarPermissionLevel struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

const (
	// Author
	CalendarPermissionLevelAuthor string = `Author`
	// Contributor
	CalendarPermissionLevelContributor string = `Contributor`
	// Custom
	CalendarPermissionLevelCustom string = `Custom`
	// Editor
	CalendarPermissionLevelEditor string = `Editor`
	// FreeBusyTimeAndSubjectAndLocation
	CalendarPermissionLevelFreeBusyTimeAndSubjectAndLocation string = `FreeBusyTimeAndSubjectAndLocation`
	// FreeBusyTimeOnly
	CalendarPermissionLevelFreeBusyTimeOnly string = `FreeBusyTimeOnly`
	// None
	CalendarPermissionLevelNone string = `None`
	// NoneditingAuthor
	CalendarPermissionLevelNoneditingAuthor string = `NoneditingAuthor`
	// Owner
	CalendarPermissionLevelOwner string = `Owner`
	// PublishingAuthor
	CalendarPermissionLevelPublishingAuthor string = `PublishingAuthor`
	// PublishingEditor
	CalendarPermissionLevelPublishingEditor string = `PublishingEditor`
	// Reviewer
	CalendarPermissionLevelReviewer string = `Reviewer`
)

func (C *CalendarPermissionLevel) SetForMarshal() {
	C.XMLName.Local = "t:CalendarPermissionLevel"
}

func (C *CalendarPermissionLevel) GetSchema() *Schema {
	return &SchemaTypes
}
