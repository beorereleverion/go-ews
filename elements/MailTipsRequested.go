package elements

// The MailTipsRequested element contains the types of mail tips requested from the service.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/mailtipsrequested
import "encoding/xml"

type MailTipsRequested struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

const (
	// Represents all available mail tips.
	MailTipsRequestedAll string = `All`
	// Represents a custom mail tip.
	MailTipsRequestedCustomMailTip string = `CustomMailTip`
	// Indicates whether delivery restrictions will prevent the sender's message from reaching the recipient.
	MailTipsRequestedDeliveryRestriction string = `DeliveryRestriction`
	// Represents the count of external members.
	MailTipsRequestedExternalMemberCount string = `ExternalMemberCount`
	// Indicates whether the recipient is invalid.
	MailTipsRequestedInvalidRecipient string = `InvalidRecipient`
	// Represents the status for a mailbox that is full.
	MailTipsRequestedMailboxFullStatus string = `MailboxFullStatus`
	// Represents the maximum message size a recipient can accept.
	MailTipsRequestedMaxMessageSize string = `MaxMessageSize`
	// Indicates whether the sender's message will be reviewed by a moderator.
	MailTipsRequestedModerationStatus string = `ModerationStatus`
	// Represents the Out of Office (OOF) message.
	MailTipsRequestedOutOfOfficeMessage string = `OutOfOfficeMessage`
	// Represents the count of all members.
	MailTipsRequestedTotalMemberCount string = `TotalMemberCount`
)

func (M *MailTipsRequested) SetForMarshal() {
	M.XMLName.Local = "m:MailTipsRequested"
}

func (M *MailTipsRequested) GetSchema() *Schema {
	return &SchemaMessages
}
