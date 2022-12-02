package elements

// The UnifiedMessagingConfiguration element contains service configuration information for the Unified Messaging service.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/unifiedmessagingconfiguration
type UnifiedMessagingConfiguration struct {
	// The PlayOnPhoneDialString element identifies the Play-on-Phone dial string.
	PlayOnPhoneDialString *PlayOnPhoneDialStringExchangeWebServices `xml:"t:PlayOnPhoneDialString"`
	// The PlayOnPhoneEnabled element indicates whether the Play-on-Phone feature is enabled.
	PlayOnPhoneEnabled *PlayOnPhoneEnabled `xml:"t:PlayOnPhoneEnabled"`
	// The UmEnabled element indicates whether Unified Messaging is enabled for an account.
	UmEnabled *UmEnabled `xml:"t:UmEnabled"`
}
