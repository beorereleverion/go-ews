package elements

// The FractionalPageItemView element describes where the paged view starts and the maximum number of items returned in a FindItem request.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/fractionalpageitemview
import "encoding/xml"

type FractionalPageItemView struct {
	XMLName xml.Name

	// Represents the denominator of the fractional offset from the start of the total number of items in the result set. This attribute is required. This attribute must represent an integral value that is greater than one.   For more information, see Remarks later in this topic.
	Denominator *string `xml:"Denominator,attr"`
	// Identifies the maximum number of results to return in the FindItem response. This attribute is optional. If this attribute is not specified, the call will return all available items.
	MaxEntriesReturned *string `xml:"MaxEntriesReturned,attr"`
	// Represents the numerator of the fractional offset from the start of the result set. This attribute is required. The numerator must be equal to or less than the denominator. This attribute must represent an integral value that is equal to or greater than zero.   For more information, see Remarks later in this topic.
	Numerator *string `xml:"Numerator,attr"`
}

func (F *FractionalPageItemView) SetForMarshal() {
	F.XMLName.Local = "m:FractionalPageItemView"
}

func (F *FractionalPageItemView) GetSchema() *Schema {
	return &SchemaMessages
}
