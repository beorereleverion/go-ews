package elements

// The ResponseCode element provides information about why the recipient is invalid.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/responsecode-invalidrecipientresponsecodetype
import "encoding/xml"

type ResponseCodeInvalidRecipientResponseCodeType struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

const (
	// Indicates that there was a problem obtaining a security token from the token server.
	ResponseCodeInvalidRecipientResponseCodeTypeCannotObtainTokenFromSTS string = `CannotObtainTokenFromSTS`
	// Indicates that the error is not specified by another error response code.
	ResponseCodeInvalidRecipientResponseCodeTypeOtherError string = `OtherError`
	// Indicates that the secure token service that is used by the specified recipient is unknown.
	ResponseCodeInvalidRecipientResponseCodeTypeRecipientOrganizationFederatedWithUnknownTokenIssuer string = `RecipientOrganizationFederatedWithUnknownTokenIssuer`
	// Indicates that a sharing relationship is not available with the organization specified in the recipient's SMTP e-mail address.
	ResponseCodeInvalidRecipientResponseCodeTypeRecipientOrganizationNotFederated string = `RecipientOrganizationNotFederated`
	// Indicates that the system administrator has set a system policy that blocks sharing with the specified recipient.
	ResponseCodeInvalidRecipientResponseCodeTypeSystemPolicyBlocksSharingWithThisRecipient string = `SystemPolicyBlocksSharingWithThisRecipient`
)

func (R *ResponseCodeInvalidRecipientResponseCodeType) SetForMarshal() {
	R.XMLName.Local = "t:ResponseCode"
}

func (R *ResponseCodeInvalidRecipientResponseCodeType) GetSchema() *Schema {
	return &SchemaTypes
}
