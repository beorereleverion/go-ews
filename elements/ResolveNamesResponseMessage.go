package elements

// The ResolveNamesResponseMessage element contains the status and result of a ResolveNames operation request.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/resolvenamesresponsemessage
import "encoding/xml"

type ResolveNamesResponseMessage struct {
	XMLName xml.Name

	// The DescriptiveLinkKey element is currently unused and is reserved for future use. It contains a value of 0.
	DescriptiveLinkKey *DescriptiveLinkKey `xml:"DescriptiveLinkKey"`
	// The MessageText element provides a text description of the status of the response.
	MessageText *MessageText `xml:"MessageText"`
	// The MessageXml element provides additional error response information.
	MessageXml *MessageXml `xml:"MessageXml"`
	// The ResolutionSet element contains an array of resolutions for an ambiguous name.
	ResolutionSet *ResolutionSet `xml:"ResolutionSet"`
	// The ResponseCode element provides status information about the request.
	ResponseCode *ResponseCode `xml:"ResponseCode"`
	// Describes the status of a ResolveNames operation response. The following values are valid for this attribute:  - Success  - Warning  - Error
	ResponseClass *string `xml:"ResponseClass,attr"`
}

func (R *ResolveNamesResponseMessage) SetForMarshal() {
	R.XMLName.Local = "m:ResolveNamesResponseMessage"
}

func (R *ResolveNamesResponseMessage) GetSchema() *Schema {
	return &SchemaMessages
}
