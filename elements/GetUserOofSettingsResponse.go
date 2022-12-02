package elements

// The GetUserOofSettingsResponse element contains the response message and the Out of Office (OOF) settings for a user.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/getuseroofsettingsresponse
type GetUserOofSettingsResponse struct {
	// The AllowExternalOof element contains a value that identifies to whom external Out of Office (OOF) messages are sent.
	AllowExternalOof *AllowExternalOof `xml:"m:AllowExternalOof"`
	// The OofSettings element contains the Out of Office (OOF) settings.
	OofSettings *OofSettings `xml:"m:OofSettings"`
	// The ResponseMessage element provides descriptive information about the response status for a single entity within a request.
	ResponseMessage *ResponseMessage `xml:"m:ResponseMessage"`
}
