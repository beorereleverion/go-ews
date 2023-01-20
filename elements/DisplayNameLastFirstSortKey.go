package elements

// The DisplayNameLastFirstSortKey element contains the sort key for a display name in last name, first name order.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/displaynamelastfirstsortkey
import "encoding/xml"

type DisplayNameLastFirstSortKey struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (D *DisplayNameLastFirstSortKey) SetForMarshal() {
	D.XMLName.Local = "t:DisplayNameLastFirstSortKey"
}

func (D *DisplayNameLastFirstSortKey) GetSchema() *Schema {
	return &SchemaTypes
}
