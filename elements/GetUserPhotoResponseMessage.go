package elements

// The GetUserPhotoResponseMessage element contains the response to a GetUserPhoto request.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/getuserphotoresponsemessage
type GetUserPhotoResponseMessage struct {
	// The DescriptiveLinkKey element is currently unused and is reserved for future use. It contains a value of 0.
	DescriptiveLinkKey *DescriptiveLinkKey `xml:"m:DescriptiveLinkKey"`
	// The HasChanged element indicates whether a user's photo has changed.
	HasChanged *HasChanged `xml:"t:HasChanged"`
	// The MessageText element provides a text description of the status of the response.
	MessageText *MessageText `xml:"m:MessageText"`
	// The MessageXml element provides additional error response information.
	MessageXml *MessageXml `xml:"m:MessageXml"`
	// The PictureData element contains the stream of picture data.
	PictureData *PictureData `xml:"m:PictureData"`
	// The ResponseCode element provides status information about the request.
	ResponseCode *ResponseCode `xml:"m:ResponseCode"`
}
