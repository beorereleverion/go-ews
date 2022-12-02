package elements

// The ReadItems element indicates whether a user has permission to read items within a Calendar folder. This element was introduced in Microsoft Exchange Server 2007 Service Pack 1 (SP1).
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/readitems-calendarpermissiontype
type ReadItemsCalendarPermissionType string

const (
	// FullDetails
	ReadItemsCalendarPermissionTypeFullDetails ReadItemsCalendarPermissionType = `FullDetails`
	// None
	ReadItemsCalendarPermissionTypeNone ReadItemsCalendarPermissionType = `None`
	// TimeAndSubjectAndLocation
	ReadItemsCalendarPermissionTypeTimeAndSubjectAndLocation ReadItemsCalendarPermissionType = `TimeAndSubjectAndLocation`
	// TimeOnly
	ReadItemsCalendarPermissionTypeTimeOnly ReadItemsCalendarPermissionType = `TimeOnly`
)
