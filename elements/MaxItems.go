package elements

// The MaxItems element specifies the maximum number of items to return in the request.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/maxitems
import "encoding/xml"

type MaxItems struct {
	XMLName xml.Name
	TEXT    int64 `xml:",chardata"`
}

func (M *MaxItems) SetForMarshal() {
	M.XMLName.Local = "m:MaxItems"
}

func (M *MaxItems) GetSchema() *Schema {
	return &SchemaMessages
}
