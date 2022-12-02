package elements

// The PendingMailTips element indicates that the mail tips in this element could not be evaluated before the server's processing timeout expired.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/pendingmailtips
type PendingMailTips string

const (
	// Represents all available mail tips.
	PendingMailTipsAll PendingMailTips = `All`
	// Represents a custom mail tip.
	PendingMailTipsCustomMailTip PendingMailTips = `CustomMailTip`
	// Indicates whether delivery restrictions will prevent the sender's message from reaching the recipient.
	PendingMailTipsDeliveryRestriction PendingMailTips = `DeliveryRestriction`
	// Represents the count of external members.
	PendingMailTipsExternalMemberCount PendingMailTips = `ExternalMemberCount`
	// Indicates whether the recipient is invalid.
	PendingMailTipsInvalidRecipient PendingMailTips = `InvalidRecipient`
	// Represents the status for a mailbox being full.
	PendingMailTipsMailboxFullStatus PendingMailTips = `MailboxFullStatus`
	// Represents the maximum message size a recipient can accept.
	PendingMailTipsMaxMessageSize PendingMailTips = `MaxMessageSize`
	// Indicates whether the sender's message will be reviewed by a moderator.
	PendingMailTipsModerationStatus PendingMailTips = `ModerationStatus`
	// Represents the Out of Office (OOF) message.
	PendingMailTipsOutOfOfficeMessage PendingMailTips = `OutOfOfficeMessage`
	// Represents the count of all members.
	PendingMailTipsTotalMemberCount PendingMailTips = `TotalMemberCount`
)
