package elements

// The CalendarFolderPermissionLevel element contains the permissions for the default Calendar folder. This element was introduced in Microsoft Exchange Server 2007 Service Pack 1 (SP1).
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/calendarfolderpermissionlevel
import "encoding/xml"

type CalendarFolderPermissionLevel struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

const (
	// The delegate user can read and create items in the Calendar folder.
	CalendarFolderPermissionLevelAuthor string = `Author`
	// The delegate user has custom access permissions to the Calendar folder.
	CalendarFolderPermissionLevelCustom string = `Custom`
	// The delegate user can read, create, and modify items in the Calendar folder.
	CalendarFolderPermissionLevelEditor string = `Editor`
	// The delegate user has no access permissions to the Calendar folder.
	CalendarFolderPermissionLevelNone string = `None`
	// The delegate user can read items in the Calendar folder.
	CalendarFolderPermissionLevelReviewer string = `Reviewer`
)

func (C *CalendarFolderPermissionLevel) SetForMarshal() {
	C.XMLName.Local = "t:CalendarFolderPermissionLevel"
}

func (C *CalendarFolderPermissionLevel) GetSchema() *Schema {
	return &SchemaTypes
}
