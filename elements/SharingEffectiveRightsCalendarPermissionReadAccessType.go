package elements

// The SharingEffectiveRights element indicates the permissions that the user has for the calendar data that is being shared.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/sharingeffectiverights-calendarpermissionreadaccesstype
import "encoding/xml"

type SharingEffectiveRightsCalendarPermissionReadAccessType struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

const (
	// Indicates that the user has permission to view all items in the calendar, including free/busy time and subject, location, and details of appointments.
	SharingEffectiveRightsCalendarPermissionReadAccessTypeFullDetails string = `FullDetails`
	// Indicates that the user does not have permission to view items in the calendar.
	SharingEffectiveRightsCalendarPermissionReadAccessTypeNone string = `None`
	// Indicates that the user has permission to view free/busy time in the calendar and the subject and location of appointments.
	SharingEffectiveRightsCalendarPermissionReadAccessTypeTimeAndSubjectAndLocation string = `TimeAndSubjectAndLocation`
	// Indicates that the user has permission to view only free/busy time in the calendar.
	SharingEffectiveRightsCalendarPermissionReadAccessTypeTimeOnly string = `TimeOnly`
)

func (S *SharingEffectiveRightsCalendarPermissionReadAccessType) SetForMarshal() {
	S.XMLName.Local = "t:SharingEffectiveRights"
}

func (S *SharingEffectiveRightsCalendarPermissionReadAccessType) GetSchema() *Schema {
	return &SchemaTypes
}
