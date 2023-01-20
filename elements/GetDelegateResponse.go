package elements

// The GetDelegateResponse element contains the status and result of a GetDelegate operation request.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/getdelegateresponse
import "encoding/xml"

type GetDelegateResponse struct {
	XMLName xml.Name

	// The DeliverMeetingRequests element defines how meeting requests are handled between the delegate and the principal. This element was introduced in Microsoft Exchange Server 2007 Service Pack 1 (SP1).
	DeliverMeetingRequests *DeliverMeetingRequests `xml:"DeliverMeetingRequests"`
	// The DescriptiveLinkKey element is currently unused and is reserved for future use. It contains a value of 0.
	DescriptiveLinkKey *DescriptiveLinkKey `xml:"DescriptiveLinkKey"`
	// The MessageText element provides a text description of the status of the response.
	MessageText *MessageText `xml:"MessageText"`
	// The MessageXml element provides additional error response information.
	MessageXml *MessageXml `xml:"MessageXml"`
	// The ResponseCode element provides status information about the request.
	ResponseCode *ResponseCode `xml:"ResponseCode"`
	// The ResponseMessages element contains the response messages for an Exchange Web Services delegate management request.
	ResponseMessages *ResponseMessagesArrayOfDelegateUserResponseMessageType `xml:"ResponseMessages"`
}

func (G *GetDelegateResponse) SetForMarshal() {
	G.XMLName.Local = "m:GetDelegateResponse"
}

func (G *GetDelegateResponse) GetSchema() *Schema {
	return &SchemaMessages
}
