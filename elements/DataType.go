package elements

// The DataType element describes the type of data that is shared by a shared folder.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/datatype
type DataType string

const (
	// Calendar
	DataTypeCalendar DataType = `Calendar`
	// Contacts
	DataTypeContacts DataType = `Contacts`
)
