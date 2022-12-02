package elements

// The RetentionPolicyType element specifies the retention policy type applied to items in a conversation.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/retentionpolicytype
type RetentionPolicyType string

const (
	// Archive
	RetentionPolicyTypeArchive RetentionPolicyType = `Archive`
	// Delete
	RetentionPolicyTypeDelete RetentionPolicyType = `Delete`
)
