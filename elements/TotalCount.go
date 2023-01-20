package elements

// The TotalCount element represents the total count of items within a given folder.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/totalcount
import "encoding/xml"

type TotalCount struct {
	XMLName xml.Name
	TEXT    int64 `xml:",chardata"`
}

func (T *TotalCount) SetForMarshal() {
	T.XMLName.Local = "t:TotalCount"
}

func (T *TotalCount) GetSchema() *Schema {
	return &SchemaTypes
}
