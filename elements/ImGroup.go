package elements

// The ImGroup element represents an instant messaging group.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/imgroup
type ImGroup struct {
	// The DisplayName element contains the display name of a new instant messaging group contact or the display name of a new instant messaging group.
	DisplayName *DisplayNameNonEmptyStringType `xml:"t:DisplayName"`
	// The ExchangeStoreId element specifies the instant messaging (IM) group identifier.
	ExchangeStoreId *ExchangeStoreId `xml:"t:ExchangeStoreId"`
	// The ExtendedProperties element specifies an array of additional properties.
	ExtendedProperties *ExtendedPropertiesNonEmptyArrayOfExtendedPropertyType `xml:"t:ExtendedProperties"`
	// The GroupType element specifies the group class of an instant messaging (IM) group.
	GroupType *GroupType `xml:"t:GroupType"`
	// The MemberCorrelationKey element specifies the identifiers of the contacts that are part of the instant messaging (IM) group.
	MemberCorrelationKey *MemberCorrelationKey `xml:"t:MemberCorrelationKey"`
	// The SmtpAddress element represents the Simple Mail Transfer Protocol (SMTP) address of an account to be used for impersonation or a Simple Mail Transfer Protocol (SMTP) recipient address of a calendar or contact sharing request.
	SmtpAddress *SmtpAddress `xml:"t:SmtpAddress"`
}
