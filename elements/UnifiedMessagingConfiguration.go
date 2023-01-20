package elements

// The UnifiedMessagingConfiguration element contains service configuration information for the Unified Messaging service.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/unifiedmessagingconfiguration
import "encoding/xml"

type UnifiedMessagingConfiguration struct {
	XMLName xml.Name

	// The PlayOnPhoneDialString element identifies the Play-on-Phone dial string.
	PlayOnPhoneDialString *PlayOnPhoneDialStringExchangeWebServices `xml:"PlayOnPhoneDialString"`
	// The PlayOnPhoneEnabled element indicates whether the Play-on-Phone feature is enabled.
	PlayOnPhoneEnabled *PlayOnPhoneEnabled `xml:"PlayOnPhoneEnabled"`
	// The UmEnabled element indicates whether Unified Messaging is enabled for an account.
	UmEnabled *UmEnabled `xml:"UmEnabled"`
}

func (U *UnifiedMessagingConfiguration) SetForMarshal() {
	U.XMLName.Local = "m:UnifiedMessagingConfiguration"
}

func (U *UnifiedMessagingConfiguration) GetSchema() *Schema {
	return &SchemaMessages
}
