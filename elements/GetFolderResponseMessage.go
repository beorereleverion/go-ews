package elements

// The GetFolderResponseMessage element contains the status and result of a single GetFolder operation request.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/getfolderresponsemessage
import "encoding/xml"

type GetFolderResponseMessage struct {
	XMLName xml.Name

	// The DescriptiveLinkKey element is currently unused and is reserved for future use. It contains a value of 0.
	DescriptiveLinkKey *DescriptiveLinkKey `xml:"DescriptiveLinkKey"`
	// The Folders element contains an array of folders that are used in folder operations.
	Folders *Folders `xml:"Folders"`
	// The MessageText element provides a text description of the status of the response.
	MessageText *MessageText `xml:"MessageText"`
	// The MessageXml element provides additional error response information.
	MessageXml *MessageXml `xml:"MessageXml"`
	// The ResponseCode element provides status information about the request.
	ResponseCode *ResponseCode `xml:"ResponseCode"`
	// Describes the status of a GetFolder operation response. The following values are valid for this attribute:  - Success  - Warning  - Error
	ResponseClass *string `xml:"ResponseClass,attr"`
}

func (G *GetFolderResponseMessage) SetForMarshal() {
	G.XMLName.Local = "m:GetFolderResponseMessage"
}

func (G *GetFolderResponseMessage) GetSchema() *Schema {
	return &SchemaMessages
}
