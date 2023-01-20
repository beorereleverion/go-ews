package elements

// The RemoveContactFromImListResponse element represents a response to a RemoveContactFromImList request.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/removecontactfromimlistresponse
import "encoding/xml"

type RemoveContactFromImListResponse struct {
	XMLName xml.Name

	// The DescriptiveLinkKey element is currently unused and is reserved for future use. It contains a value of 0.
	DescriptiveLinkKey *DescriptiveLinkKey `xml:"DescriptiveLinkKey"`
	// The MessageText element provides a text description of the status of the response.
	MessageText *MessageText `xml:"MessageText"`
	// The MessageXml element provides additional error response information.
	MessageXml *MessageXml `xml:"MessageXml"`
	// The ResponseCode element provides status information about the request.
	ResponseCode *ResponseCode `xml:"ResponseCode"`
}

func (R *RemoveContactFromImListResponse) SetForMarshal() {
	R.XMLName.Local = "m:RemoveContactFromImListResponse"
}

func (R *RemoveContactFromImListResponse) GetSchema() *Schema {
	return &SchemaMessages
}
