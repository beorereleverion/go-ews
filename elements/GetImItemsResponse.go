package elements

// The GetImItemsResponse element defines a response to a GetImItems request.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/getimitemsresponse
import "encoding/xml"

type GetImItemsResponse struct {
	XMLName xml.Name

	// The DescriptiveLinkKey element is currently unused and is reserved for future use. It contains a value of 0.
	DescriptiveLinkKey *DescriptiveLinkKey `xml:"DescriptiveLinkKey"`
	// The ImItemList element contains a list of instant messaging groups and instant messaging contacts.
	ImItemList *ImItemList `xml:"ImItemList"`
	// The MessageText element provides a text description of the status of the response.
	MessageText *MessageText `xml:"MessageText"`
	// The MessageXml element provides additional error response information.
	MessageXml *MessageXml `xml:"MessageXml"`
	// The ResponseCode element provides status information about the request.
	ResponseCode *ResponseCode `xml:"ResponseCode"`
}

func (G *GetImItemsResponse) SetForMarshal() {
	G.XMLName.Local = "m:GetImItemsResponse"
}

func (G *GetImItemsResponse) GetSchema() *Schema {
	return &SchemaMessages
}
