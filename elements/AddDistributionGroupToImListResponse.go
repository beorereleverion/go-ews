package elements

// The AddDistributionGroupToImListResponse element defines a response to a AddDistributionGroupToImList request.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/adddistributiongrouptoimlistresponse
import "encoding/xml"

type AddDistributionGroupToImListResponse struct {
	XMLName xml.Name

	// The DescriptiveLinkKey element is currently unused and is reserved for future use. It contains a value of 0.
	DescriptiveLinkKey *DescriptiveLinkKey `xml:"DescriptiveLinkKey"`
	// The ImGroup element represents an instant messaging group.
	ImGroup *ImGroup `xml:"ImGroup"`
	// The MessageText element provides a text description of the status of the response.
	MessageText *MessageText `xml:"MessageText"`
	// The MessageXml element provides additional error response information.
	MessageXml *MessageXml `xml:"MessageXml"`
	// The ResponseCode element provides status information about the request.
	ResponseCode *ResponseCode `xml:"ResponseCode"`
}

func (A *AddDistributionGroupToImListResponse) SetForMarshal() {
	A.XMLName.Local = "m:AddDistributionGroupToImListResponse"
}

func (A *AddDistributionGroupToImListResponse) GetSchema() *Schema {
	return &SchemaMessages
}
