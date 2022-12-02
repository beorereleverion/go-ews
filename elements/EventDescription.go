package elements

// The EventDescription element
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/eventdescription
type EventDescription string

const (
	//
	EventDescriptionApprovedModeration EventDescription = `ApprovedModeration`
	//
	EventDescriptionDelayedAfterTransferToPartnerOrg EventDescription = `DelayedAfterTransferToPartnerOrg`
	//
	EventDescriptionDelivered EventDescription = `Delivered`
	//
	EventDescriptionExpanded EventDescription = `Expanded`
	//
	EventDescriptionFailedGeneral EventDescription = `FailedGeneral`
	//
	EventDescriptionFailedModeration EventDescription = `FailedModeration`
	//
	EventDescriptionFailedTransportRules EventDescription = `FailedTransportRules`
	//
	EventDescriptionForwarded EventDescription = `Forwarded`
	//
	EventDescriptionMessageDefer EventDescription = `MessageDefer`
	//
	EventDescriptionMovedToFolderByInboxRule EventDescription = `MovedToFolderByInboxRule`
	//
	EventDescriptionNotRead EventDescription = `NotRead`
	//
	EventDescriptionPending EventDescription = `Pending`
	//
	EventDescriptionPendingModeration EventDescription = `PendingModeration`
	//
	EventDescriptionQueueRetry EventDescription = `QueueRetry`
	//
	EventDescriptionQueueRetryNoRetryTime EventDescription = `QueueRetryNoRetryTime`
	//
	EventDescriptionRead EventDescription = `Read`
	//
	EventDescriptionResolved EventDescription = `Resolved`
	//
	EventDescriptionRulesCc EventDescription = `RulesCc`
	//
	EventDescriptionSmtpReceive EventDescription = `SmtpReceive`
	//
	EventDescriptionSmtpSend EventDescription = `SmtpSend`
	//
	EventDescriptionSmtpSendCrossForest EventDescription = `SmtpSendCrossForest`
	//
	EventDescriptionSmtpSendCrossSite EventDescription = `SmtpSendCrossSite`
	//
	EventDescriptionSubmitted EventDescription = `Submitted`
	//
	EventDescriptionTransferredToForeignOrg EventDescription = `TransferredToForeignOrg`
	//
	EventDescriptionTransferredToLegacyExchangeServer EventDescription = `TransferredToLegacyExchangeServer`
	//
	EventDescriptionTransferredToPartnerOrg EventDescription = `TransferredToPartnerOrg`
)
