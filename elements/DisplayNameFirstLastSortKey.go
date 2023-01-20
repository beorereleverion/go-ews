package elements

// The DisplayNameFirstLastSortKey element contains the sort key for a display name in first name, last name order.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/displaynamefirstlastsortkey
import "encoding/xml"

type DisplayNameFirstLastSortKey struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (D *DisplayNameFirstLastSortKey) SetForMarshal() {
	D.XMLName.Local = "t:DisplayNameFirstLastSortKey"
}

func (D *DisplayNameFirstLastSortKey) GetSchema() *Schema {
	return &SchemaTypes
}
