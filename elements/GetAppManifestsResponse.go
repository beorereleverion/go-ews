package elements

// The GetAppManifestsResponse element defines the response for a GetAppManifests operation request.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/getappmanifestsresponse
import "encoding/xml"

type GetAppManifestsResponse struct {
	XMLName xml.Name

	// The Apps element contains information about all the XML manifest files for apps installed in a mailbox.
	Apps *Apps `xml:"Apps"`
	// The Manifests element contains a collection of base64-encoded app manifests that are installed for the email account.
	Manifests *Manifests `xml:"Manifests"`
	// The ResponseCode element provides status information about the request.
	ResponseCode *ResponseCode `xml:"ResponseCode"`
}

func (G *GetAppManifestsResponse) SetForMarshal() {
	G.XMLName.Local = "m:GetAppManifestsResponse"
}

func (G *GetAppManifestsResponse) GetSchema() *Schema {
	return &SchemaMessages
}
