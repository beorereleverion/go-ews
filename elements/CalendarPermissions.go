package elements

// The CalendarPermissions element contains an array of calendar permissions for a folder. This element was introduced in Microsoft Exchange Server 2007 Service Pack 1 (SP1).
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/calendarpermissions
type CalendarPermissions struct {
	// The CalendarPermission element defines the access that a user has to a Calendar folder. This element was introduced in Microsoft Exchange Server 2007 Service Pack 1 (SP1).
	CalendarPermission *CalendarPermission `xml:"t:CalendarPermission"`
}
