package elements

// The GetUserRetentionPolicyTagsResponse element contains the response to a GetRetentionPolicyTags request.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/getuserretentionpolicytagsresponse
import "encoding/xml"

type GetUserRetentionPolicyTagsResponse struct {
	XMLName xml.Name

	// The DescriptiveLinkKey element is currently unused and is reserved for future use. It contains a value of 0.
	DescriptiveLinkKey *DescriptiveLinkKey `xml:"DescriptiveLinkKey"`
	// The MessageText element provides a text description of the status of the response.
	MessageText *MessageText `xml:"MessageText"`
	// The MessageXml element provides additional error response information.
	MessageXml *MessageXml `xml:"MessageXml"`
	// The ResponseCode element provides status information about the request.
	ResponseCode *ResponseCode `xml:"ResponseCode"`
	// The RetentionPolicyTags element contains a list of retention tags returned in the response of the GetUserRetentionPolicyTags WSDL operation.
	RetentionPolicyTags *RetentionPolicyTags `xml:"RetentionPolicyTags"`
}

func (G *GetUserRetentionPolicyTagsResponse) SetForMarshal() {
	G.XMLName.Local = "m:GetUserRetentionPolicyTagsResponse"
}

func (G *GetUserRetentionPolicyTagsResponse) GetSchema() *Schema {
	return &SchemaMessages
}
