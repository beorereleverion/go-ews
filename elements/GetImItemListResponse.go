package elements

// The GetImItemListResponse element defines a response to a GetImItemList request.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/getimitemlistresponse
import "encoding/xml"

type GetImItemListResponse struct {
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

func (G *GetImItemListResponse) SetForMarshal() {
	G.XMLName.Local = "m:GetImItemListResponse"
}

func (G *GetImItemListResponse) GetSchema() *Schema {
	return &SchemaMessages
}
