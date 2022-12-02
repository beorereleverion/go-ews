package elements

// The MailTipsRequested element contains the types of mail tips requested from the service.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/mailtipsrequested
type MailTipsRequested string

const (
	// Represents all available mail tips.
	MailTipsRequestedAll MailTipsRequested = `All`
	// Represents a custom mail tip.
	MailTipsRequestedCustomMailTip MailTipsRequested = `CustomMailTip`
	// Indicates whether delivery restrictions will prevent the sender's message from reaching the recipient.
	MailTipsRequestedDeliveryRestriction MailTipsRequested = `DeliveryRestriction`
	// Represents the count of external members.
	MailTipsRequestedExternalMemberCount MailTipsRequested = `ExternalMemberCount`
	// Indicates whether the recipient is invalid.
	MailTipsRequestedInvalidRecipient MailTipsRequested = `InvalidRecipient`
	// Represents the status for a mailbox that is full.
	MailTipsRequestedMailboxFullStatus MailTipsRequested = `MailboxFullStatus`
	// Represents the maximum message size a recipient can accept.
	MailTipsRequestedMaxMessageSize MailTipsRequested = `MaxMessageSize`
	// Indicates whether the sender's message will be reviewed by a moderator.
	MailTipsRequestedModerationStatus MailTipsRequested = `ModerationStatus`
	// Represents the Out of Office (OOF) message.
	MailTipsRequestedOutOfOfficeMessage MailTipsRequested = `OutOfOfficeMessage`
	// Represents the count of all members.
	MailTipsRequestedTotalMemberCount MailTipsRequested = `TotalMemberCount`
)
