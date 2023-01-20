package elements

// The ReadItems element indicates whether a user has permission to read items within a Calendar folder. This element was introduced in Microsoft Exchange Server 2007 Service Pack 1 (SP1).
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/readitems-calendarpermissiontype
import "encoding/xml"

type ReadItemsCalendarPermissionType struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

const (
	// FullDetails
	ReadItemsCalendarPermissionTypeFullDetails string = `FullDetails`
	// None
	ReadItemsCalendarPermissionTypeNone string = `None`
	// TimeAndSubjectAndLocation
	ReadItemsCalendarPermissionTypeTimeAndSubjectAndLocation string = `TimeAndSubjectAndLocation`
	// TimeOnly
	ReadItemsCalendarPermissionTypeTimeOnly string = `TimeOnly`
)

func (R *ReadItemsCalendarPermissionType) SetForMarshal() {
	R.XMLName.Local = "t:ReadItems"
}

func (R *ReadItemsCalendarPermissionType) GetSchema() *Schema {
	return &SchemaTypes
}
