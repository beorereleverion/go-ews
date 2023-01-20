package elements

// The EventDescription element
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/eventdescription
import "encoding/xml"

type EventDescription struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

const (
	//
	EventDescriptionApprovedModeration string = `ApprovedModeration`
	//
	EventDescriptionDelayedAfterTransferToPartnerOrg string = `DelayedAfterTransferToPartnerOrg`
	//
	EventDescriptionDelivered string = `Delivered`
	//
	EventDescriptionExpanded string = `Expanded`
	//
	EventDescriptionFailedGeneral string = `FailedGeneral`
	//
	EventDescriptionFailedModeration string = `FailedModeration`
	//
	EventDescriptionFailedTransportRules string = `FailedTransportRules`
	//
	EventDescriptionForwarded string = `Forwarded`
	//
	EventDescriptionMessageDefer string = `MessageDefer`
	//
	EventDescriptionMovedToFolderByInboxRule string = `MovedToFolderByInboxRule`
	//
	EventDescriptionNotRead string = `NotRead`
	//
	EventDescriptionPending string = `Pending`
	//
	EventDescriptionPendingModeration string = `PendingModeration`
	//
	EventDescriptionQueueRetry string = `QueueRetry`
	//
	EventDescriptionQueueRetryNoRetryTime string = `QueueRetryNoRetryTime`
	//
	EventDescriptionRead string = `Read`
	//
	EventDescriptionResolved string = `Resolved`
	//
	EventDescriptionRulesCc string = `RulesCc`
	//
	EventDescriptionSmtpReceive string = `SmtpReceive`
	//
	EventDescriptionSmtpSend string = `SmtpSend`
	//
	EventDescriptionSmtpSendCrossForest string = `SmtpSendCrossForest`
	//
	EventDescriptionSmtpSendCrossSite string = `SmtpSendCrossSite`
	//
	EventDescriptionSubmitted string = `Submitted`
	//
	EventDescriptionTransferredToForeignOrg string = `TransferredToForeignOrg`
	//
	EventDescriptionTransferredToLegacyExchangeServer string = `TransferredToLegacyExchangeServer`
	//
	EventDescriptionTransferredToPartnerOrg string = `TransferredToPartnerOrg`
)

func (E *EventDescription) SetForMarshal() {
	E.XMLName.Local = "t:EventDescription"
}

func (E *EventDescription) GetSchema() *Schema {
	return &SchemaTypes
}
