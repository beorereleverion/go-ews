package elements

// The FindMessageTrackingReportResponse element contains the status and result of a single FindMessageTrackingReport operation request.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/findmessagetrackingreportresponse
import "encoding/xml"

type FindMessageTrackingReportResponse struct {
	XMLName xml.Name

	// The DescriptiveLinkKey element is currently unused and is reserved for future use. It contains a value of 0.
	DescriptiveLinkKey *DescriptiveLinkKey `xml:"DescriptiveLinkKey"`
	// The Diagnostics element provides timing and performance information that is used for reporting in a DataCenter.
	Diagnostics *Diagnostics `xml:"Diagnostics"`
	// The Errors element contains a property bag to store errors that are returned through the Web service.
	Errors *Errors `xml:"Errors"`
	// The ExecutedSearchScope element contains the scope of the search that was performed to get the search results.
	ExecutedSearchScope *ExecutedSearchScope `xml:"ExecutedSearchScope"`
	// The MessageText element provides a text description of the status of the response.
	MessageText *MessageText `xml:"MessageText"`
	// The MessageTrackingSearchResults element contains a list of records that match the search criteria.
	MessageTrackingSearchResults *MessageTrackingSearchResults `xml:"MessageTrackingSearchResults"`
	// The MessageXml element provides additional error response information.
	MessageXml *MessageXml `xml:"MessageXml"`
	// The Properties element contains a list of one or more tracking properties.
	Properties *PropertiesArrayOfTrackingPropertiesType `xml:"Properties"`
	// The ResponseCode element provides status information about the request.
	ResponseCode *ResponseCode `xml:"ResponseCode"`
	// Describes the status of the response. The following values are valid for this attribute:  - Success  - Warning  - Error
	ResponseClass *string `xml:"ResponseClass,attr"`
}

const (
	// Describes a request that is fulfilled.
	FindMessageTrackingReportResponseSuccess = `Success`
	// Describes a request that was not processed. A warning may be returned if an error occurred while an item in the request was processing and subsequent items could not be processed. The following are examples of sources of warnings:  - The Exchange store is offline during the batch.  - Active Directory Domain Services (AD DS) is offline.  - Mailboxes were moved.  - The message database (MDB) is offline.  - A password is expired.  - A quota has been exceeded.
	FindMessageTrackingReportResponseWarning = `Warning`
	// Describes a request that cannot be fulfilled. The following are examples of sources of errors:  - Invalid attributes or elements  - Attributes or elements that are out of range  - An unknown tag  - An attribute or element that is not valid in the context  - An unauthorized access attempt by any client  - A server-side failure in response to a valid client-side call    Information about the error can be found in the ResponseCode and MessageText elements.
	FindMessageTrackingReportResponseError = `Error`
)

func (F *FindMessageTrackingReportResponse) SetForMarshal() {
	F.XMLName.Local = "m:FindMessageTrackingReportResponse"
}

func (F *FindMessageTrackingReportResponse) GetSchema() *Schema {
	return &SchemaMessages
}
