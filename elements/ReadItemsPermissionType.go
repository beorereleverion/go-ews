package elements

// The ReadItems element indicates whether a user has permission to read items within a folder. This element was introduced in Microsoft Exchange Server 2007 Service Pack 1 (SP1).
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/readitems-permissiontype
import "encoding/xml"

type ReadItemsPermissionType struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

const (
	// FullDetails
	ReadItemsPermissionTypeFullDetails string = `FullDetails`
	// None
	ReadItemsPermissionTypeNone string = `None`
)

func (R *ReadItemsPermissionType) SetForMarshal() {
	R.XMLName.Local = "t:ReadItems"
}

func (R *ReadItemsPermissionType) GetSchema() *Schema {
	return &SchemaTypes
}
