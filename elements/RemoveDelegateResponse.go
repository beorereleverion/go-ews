package elements

// The RemoveDelegateResponse element contains the status and result of a RemoveDelegate operation request.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/removedelegateresponse
import "encoding/xml"

type RemoveDelegateResponse struct {
	XMLName xml.Name

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

func (R *RemoveDelegateResponse) SetForMarshal() {
	R.XMLName.Local = "m:RemoveDelegateResponse"
}

func (R *RemoveDelegateResponse) GetSchema() *Schema {
	return &SchemaMessages
}
