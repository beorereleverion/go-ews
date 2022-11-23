package elements

// The GetFolderResponseMessage element contains the status and result of a single GetFolder operation request.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/getfolderresponsemessage
type GetFolderResponseMessage struct {
	ResponseClass      string `xml:"ResponseClass,attr"`
	MessageText        string `xml:"MessageText"`
	ResponseCode       string `xml:"ResponseCode"`
	DescriptiveLinkKey string `xml:"DescriptiveLinkKey"`
	MessageXml         string `xml:"MessageXml"`
	Folders            string `xml:"Folders"`
}
