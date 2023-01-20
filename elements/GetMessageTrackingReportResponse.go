package elements

// The GetMessageTrackingReportResponse element contains the response for the GetMessageTrackingReport operation.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/getmessagetrackingreportresponse
import "encoding/xml"

type GetMessageTrackingReportResponse struct {
	XMLName xml.Name

	// The DescriptiveLinkKey element is currently unused and is reserved for future use. It contains a value of 0.
	DescriptiveLinkKey *DescriptiveLinkKey `xml:"DescriptiveLinkKey"`
	// The Diagnostics element provides timing and performance information that is used for reporting in a DataCenter.
	Diagnostics *Diagnostics `xml:"Diagnostics"`
	// The Errors element contains a property bag to store errors that are returned through the Web service.
	Errors *Errors `xml:"Errors"`
	// The MessageText element provides a text description of the status of the response.
	MessageText *MessageText `xml:"MessageText"`
	// The MessageTrackingReport element contains a single message that is returned in a GetMessageTrackingReport operation.
	MessageTrackingReport *MessageTrackingReport `xml:"MessageTrackingReport"`
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
	GetMessageTrackingReportResponseSuccess = `Success`
	// Describes a request that was not processed. A warning may be returned if an error occurred while an item in the request was processing and subsequent items could not be processed. The following are examples of sources of warnings:  - The Exchange store is offline during the batch.  - Active Directory Domain Serves (AD DS) is offline.  - Mailboxes were moved.  - The message database (MDB) is offline.  - A password is expired.  - A quota has been exceeded.
	GetMessageTrackingReportResponseWarning = `Warning`
	// Describes a request that cannot be fulfilled. The following are examples of sources of errors:    Invalid attributes or elements  - Attributes or elements that are out of range  - An unknown tag  - An attribute or element that is not valid in the context  - An unauthorized access attempt by any client  - A server-side failure in response to a valid client-side call    Information about the error can be found in the ResponseCode and MessageText elements.
	GetMessageTrackingReportResponseError = `Error`
)

func (G *GetMessageTrackingReportResponse) SetForMarshal() {
	G.XMLName.Local = "m:GetMessageTrackingReportResponse"
}

func (G *GetMessageTrackingReportResponse) GetSchema() *Schema {
	return &SchemaMessages
}
