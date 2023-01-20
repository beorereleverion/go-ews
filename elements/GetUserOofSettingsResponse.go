package elements

// The GetUserOofSettingsResponse element contains the response message and the Out of Office (OOF) settings for a user.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/getuseroofsettingsresponse
import "encoding/xml"

type GetUserOofSettingsResponse struct {
	XMLName xml.Name

	// The AllowExternalOof element contains a value that identifies to whom external Out of Office (OOF) messages are sent.
	AllowExternalOof *AllowExternalOof `xml:"AllowExternalOof"`
	// The OofSettings element contains the Out of Office (OOF) settings.
	OofSettings *OofSettings `xml:"OofSettings"`
	// The ResponseMessage element provides descriptive information about the response status for a single entity within a request.
	ResponseMessage *ResponseMessage `xml:"ResponseMessage"`
}

func (G *GetUserOofSettingsResponse) SetForMarshal() {
	G.XMLName.Local = "m:GetUserOofSettingsResponse"
}

func (G *GetUserOofSettingsResponse) GetSchema() *Schema {
	return &SchemaMessages
}
