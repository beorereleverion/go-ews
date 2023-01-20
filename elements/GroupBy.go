package elements

// The GroupBy element specifies an arbitrary grouping for FindItem queries.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/groupby
import "encoding/xml"

type GroupBy struct {
	XMLName xml.Name

	// The AggregateOn element represents the property that is used to determine the order of grouped items for a grouped FindItem result set.
	AggregateOn *AggregateOn `xml:"AggregateOn"`
	// The ExtendedFieldURI element identifies an extended MAPI property.
	ExtendedFieldURI *ExtendedFieldURI `xml:"ExtendedFieldURI"`
	// The FieldURI element identifies frequently referenced properties by URI.
	FieldURI *FieldURI `xml:"FieldURI"`
	// The IndexedFieldURI element identifies individual members of a dictionary.
	IndexedFieldURI *IndexedFieldURI `xml:"IndexedFieldURI"`
	// Determines the order of the groups in the grouped item array that is returned in the response. This attribute is of type SortDirectionType.
	Order *string `xml:"Order,attr"`
}

const (
	// The groups are ordered in ascending order.
	GroupByAscending = `Ascending`
	// The groups are ordered in descending order.
	GroupByDescending = `Descending`
)

func (G *GroupBy) SetForMarshal() {
	G.XMLName.Local = "m:GroupBy"
}

func (G *GroupBy) GetSchema() *Schema {
	return &SchemaMessages
}
