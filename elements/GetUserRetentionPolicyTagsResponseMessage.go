package elements

// The GetUserRetentionPolicyTagsResponseMessage element specifies the response message for a GetUserRetentionPolicyTags request.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/getuserretentionpolicytagsresponsemessage
import "encoding/xml"

type GetUserRetentionPolicyTagsResponseMessage struct {
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

func (G *GetUserRetentionPolicyTagsResponseMessage) SetForMarshal() {
	G.XMLName.Local = "m:GetUserRetentionPolicyTagsResponseMessage"
}

func (G *GetUserRetentionPolicyTagsResponseMessage) GetSchema() *Schema {
	return &SchemaMessages
}
