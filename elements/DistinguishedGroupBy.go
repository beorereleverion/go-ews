package elements

// The DistinguishedGroupBy element provides standard groupings for FindItem queries.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/distinguishedgroupby
import "encoding/xml"

type DistinguishedGroupBy struct {
	XMLName xml.Name

	// The StandardGroupBy element represents the standard grouping and aggregating mechanisms for a grouped FindItem operation.
	StandardGroupBy *StandardGroupBy `xml:"StandardGroupBy"`
}

func (D *DistinguishedGroupBy) SetForMarshal() {
	D.XMLName.Local = "m:DistinguishedGroupBy"
}

func (D *DistinguishedGroupBy) GetSchema() *Schema {
	return &SchemaMessages
}
