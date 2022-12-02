package elements

// The ResponseCode element provides information about why the recipient is invalid.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/responsecode-invalidrecipientresponsecodetype
type ResponseCodeInvalidRecipientResponseCodeType string

const (
	// Indicates that there was a problem obtaining a security token from the token server.
	ResponseCodeInvalidRecipientResponseCodeTypeCannotObtainTokenFromSTS ResponseCodeInvalidRecipientResponseCodeType = `CannotObtainTokenFromSTS`
	// Indicates that the error is not specified by another error response code.
	ResponseCodeInvalidRecipientResponseCodeTypeOtherError ResponseCodeInvalidRecipientResponseCodeType = `OtherError`
	// Indicates that the secure token service that is used by the specified recipient is unknown.
	ResponseCodeInvalidRecipientResponseCodeTypeRecipientOrganizationFederatedWithUnknownTokenIssuer ResponseCodeInvalidRecipientResponseCodeType = `RecipientOrganizationFederatedWithUnknownTokenIssuer`
	// Indicates that a sharing relationship is not available with the organization specified in the recipient's SMTP e-mail address.
	ResponseCodeInvalidRecipientResponseCodeTypeRecipientOrganizationNotFederated ResponseCodeInvalidRecipientResponseCodeType = `RecipientOrganizationNotFederated`
	// Indicates that the system administrator has set a system policy that blocks sharing with the specified recipient.
	ResponseCodeInvalidRecipientResponseCodeTypeSystemPolicyBlocksSharingWithThisRecipient ResponseCodeInvalidRecipientResponseCodeType = `SystemPolicyBlocksSharingWithThisRecipient`
)
