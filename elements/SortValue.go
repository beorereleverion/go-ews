package elements

// The SortValue element specifies a value used for sorting.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/sortvalue
import "encoding/xml"

type SortValue struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (S *SortValue) SetForMarshal() {
	S.XMLName.Local = "t:SortValue"
}

func (S *SortValue) GetSchema() *Schema {
	return &SchemaTypes
}
