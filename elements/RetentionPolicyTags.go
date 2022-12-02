package elements

// The RetentionPolicyTags element contains a list of retention tags returned in the response of the GetUserRetentionPolicyTags WSDL operation.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/retentionpolicytags
type RetentionPolicyTags struct {
	// The RetentionPolicyTag element specifies the retention policy for a mailbox item.
	RetentionPolicyTag *RetentionPolicyTag `xml:"t:RetentionPolicyTag"`
}
