package elements

// The PendingMailTips element indicates that the mail tips in this element could not be evaluated before the server's processing timeout expired.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/pendingmailtips
import "encoding/xml"

type PendingMailTips struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

const (
	// Represents all available mail tips.
	PendingMailTipsAll string = `All`
	// Represents a custom mail tip.
	PendingMailTipsCustomMailTip string = `CustomMailTip`
	// Indicates whether delivery restrictions will prevent the sender's message from reaching the recipient.
	PendingMailTipsDeliveryRestriction string = `DeliveryRestriction`
	// Represents the count of external members.
	PendingMailTipsExternalMemberCount string = `ExternalMemberCount`
	// Indicates whether the recipient is invalid.
	PendingMailTipsInvalidRecipient string = `InvalidRecipient`
	// Represents the status for a mailbox being full.
	PendingMailTipsMailboxFullStatus string = `MailboxFullStatus`
	// Represents the maximum message size a recipient can accept.
	PendingMailTipsMaxMessageSize string = `MaxMessageSize`
	// Indicates whether the sender's message will be reviewed by a moderator.
	PendingMailTipsModerationStatus string = `ModerationStatus`
	// Represents the Out of Office (OOF) message.
	PendingMailTipsOutOfOfficeMessage string = `OutOfOfficeMessage`
	// Represents the count of all members.
	PendingMailTipsTotalMemberCount string = `TotalMemberCount`
)

func (P *PendingMailTips) SetForMarshal() {
	P.XMLName.Local = "t:PendingMailTips"
}

func (P *PendingMailTips) GetSchema() *Schema {
	return &SchemaTypes
}
