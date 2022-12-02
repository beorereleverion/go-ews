package elements

// The GetAppManifestsResponse element defines the response for a GetAppManifests operation request.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/getappmanifestsresponse
type GetAppManifestsResponse struct {
	// The Apps element contains information about all the XML manifest files for apps installed in a mailbox.
	Apps *Apps `xml:"m:Apps"`
	// The Manifests element contains a collection of base64-encoded app manifests that are installed for the email account.
	Manifests *Manifests `xml:"m:Manifests"`
	// The ResponseCode element provides status information about the request.
	ResponseCode *ResponseCode `xml:"m:ResponseCode"`
}
