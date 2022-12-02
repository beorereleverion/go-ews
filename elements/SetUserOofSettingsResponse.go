package elements

// The SetUserOofSettingsResponse element contains the result of a SetUserOofSettingsRequest message attempt.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/setuseroofsettingsresponse
type SetUserOofSettingsResponse struct {
	// The ResponseMessage element provides descriptive information about the response status for a single entity within a request.
	ResponseMessage *ResponseMessage `xml:"m:ResponseMessage"`
}
