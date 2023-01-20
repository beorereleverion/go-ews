package elements

// The MessageTrackingSearchResults element contains a list of records that match the search criteria.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/messagetrackingsearchresults
import "encoding/xml"

type MessageTrackingSearchResults struct {
	XMLName xml.Name

	// The FindMessageTrackingReportResponse element contains the status and result of a single FindMessageTrackingReport operation request.
	FindMessageTrackingReportResponse *FindMessageTrackingReportResponse `xml:"FindMessageTrackingReportResponse"`
	// The MessageTrackingSearchResult element contains a single message result for a FindMessageTrackingReportResponse element.
	MessageTrackingSearchResult *MessageTrackingSearchResult `xml:"MessageTrackingSearchResult"`
}

func (M *MessageTrackingSearchResults) SetForMarshal() {
	M.XMLName.Local = "m:MessageTrackingSearchResults"
}

func (M *MessageTrackingSearchResults) GetSchema() *Schema {
	return &SchemaMessages
}
