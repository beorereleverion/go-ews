package elements

// The SetUserOofSettingsResponse element contains the result of a SetUserOofSettingsRequest message attempt.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/setuseroofsettingsresponse
import "encoding/xml"

type SetUserOofSettingsResponse struct {
	XMLName xml.Name

	// The ResponseMessage element provides descriptive information about the response status for a single entity within a request.
	ResponseMessage *ResponseMessage `xml:"ResponseMessage"`
}

func (S *SetUserOofSettingsResponse) SetForMarshal() {
	S.XMLName.Local = "m:SetUserOofSettingsResponse"
}

func (S *SetUserOofSettingsResponse) GetSchema() *Schema {
	return &SchemaMessages
}
