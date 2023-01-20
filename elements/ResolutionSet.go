package elements

// The ResolutionSet element contains an array of resolutions for an ambiguous name.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/resolutionset
import "encoding/xml"

type ResolutionSet struct {
	XMLName xml.Name

	// The Resolution element contains a single resolved entity.
	Resolution *Resolution `xml:"Resolution"`
	// Represents the next denominator to use for the next request when you are using fraction page views.
	AbsoluteDenominator *string `xml:"AbsoluteDenominator,attr"`
	// This attribute will be true if the current results contain the last item in the query, so that additional paging is not needed.
	IncludesLastItemInRange *string `xml:"IncludesLastItemInRange,attr"`
	// Represents the next index that should be used for the next request when you are using an indexed page view.
	IndexedPagingOffset *string `xml:"IndexedPagingOffset,attr"`
	// Represents the new numerator value to use for the next request when you are using fraction page views.
	NumeratorOffset *string `xml:"NumeratorOffset,attr"`
	// Represents the total number of items in the view.
	TotalItemsInView *string `xml:"TotalItemsInView,attr"`
}

func (R *ResolutionSet) SetForMarshal() {
	R.XMLName.Local = "m:ResolutionSet"
}

func (R *ResolutionSet) GetSchema() *Schema {
	return &SchemaMessages
}
