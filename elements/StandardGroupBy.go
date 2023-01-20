package elements

// The StandardGroupBy element represents the standard grouping and aggregating mechanisms for a grouped FindItem operation.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/standardgroupby
import "encoding/xml"

type StandardGroupBy struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (S *StandardGroupBy) SetForMarshal() {
	S.XMLName.Local = "t:StandardGroupBy"
}

func (S *StandardGroupBy) GetSchema() *Schema {
	return &SchemaTypes
}
