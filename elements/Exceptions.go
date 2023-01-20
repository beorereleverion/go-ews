package elements

// The Exceptions element identifies the exceptions that represent all the available rule exception conditions for an Inbox rule.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/exceptions
import "encoding/xml"

type Exceptions struct {
	XMLName xml.Name

	// The Categories element contains a collection of strings that identify the categories to which an item in the mailbox belongs.
	Categories *Categories `xml:"Categories"`
	// The ContainsBodyStrings element indicates the strings that must appear in the body of incoming messages in order for the condition or exception to apply.
	ContainsBodyStrings *ContainsBodyStrings `xml:"ContainsBodyStrings"`
	// The ContainsHeaderStrings element indicates the strings that must appear in the headers of incoming messages in order for the condition or exception to apply.
	ContainsHeaderStrings *ContainsHeaderStrings `xml:"ContainsHeaderStrings"`
	// The ContainsRecipientStrings element indicates the strings that must appear in either the ToRecipients or CcRecipients properties of incoming messages in order for the condition or exception to apply.
	ContainsRecipientStrings *ContainsRecipientStrings `xml:"ContainsRecipientStrings"`
	// The ContainsSenderStrings element indicates the strings that must appear in the From property of incoming messages in order for the condition or exception to apply.
	ContainsSenderStrings *ContainsSenderStrings `xml:"ContainsSenderStrings"`
	// The ContainsSubjectOrBodyStrings element indicates the strings that must appear in either the body or the subject of incoming messages in order for the condition or exception to apply.
	ContainsSubjectOrBodyStrings *ContainsSubjectOrBodyStrings `xml:"ContainsSubjectOrBodyStrings"`
	// The ContainsSubjectStrings element indicates the strings that must appear in the subject of incoming messages in order for the condition or exception to apply.
	ContainsSubjectStrings *ContainsSubjectStrings `xml:"ContainsSubjectStrings"`
	// The FlaggedForAction element specifies the flag for action value that must appear on incoming messages in order for the condition or exception to apply.
	FlaggedForAction *FlaggedForAction `xml:"FlaggedForAction"`
	// The FromAddresses element indicates the e-mail addresses from which incoming messages must be sent in order for the condition or exception to apply.
	FromAddresses *FromAddresses `xml:"FromAddresses"`
	// The FromConnectedAccounts element represents the e-mail account names from which incoming messages have to have been aggregated in order for the condition or exception to apply.
	FromConnectedAccounts *FromConnectedAccounts `xml:"FromConnectedAccounts"`
	// The HasAttachments element represents a property that is set to true if an item has at least one visible attachment or if a conversation contains at least one item that has an attachment. This property is read-only.
	HasAttachments *HasAttachments `xml:"HasAttachments"`
	// The Importance element describes the importance of an item or the aggregated importance of all items in a conversation in the current folder.
	Importance *Importance `xml:"Importance"`
	// The IsApprovalRequest element indicates whether incoming messages must be approval requests in order for the condition or exception to apply.
	IsApprovalRequest *IsApprovalRequest `xml:"IsApprovalRequest"`
	// The IsAutomaticForward element indicates whether incoming messages must be automatic forwards in order for the condition or exception to apply.
	IsAutomaticForward *IsAutomaticForward `xml:"IsAutomaticForward"`
	// The IsAutomaticReply element indicates whether incoming messages must be automatic replies in order for the condition or exception to apply.
	IsAutomaticReply *IsAutomaticReply `xml:"IsAutomaticReply"`
	// The IsEncrypted element indicates whether incoming messages must be S/MIME encrypted in order for the condition or exception to apply.
	IsEncrypted *IsEncrypted `xml:"IsEncrypted"`
	// The IsMeetngRequest element indicates whether incoming messages must be a meeting request in order for the condition or exception to apply.
	IsMeetingRequest *IsMeetingRequest `xml:"IsMeetingRequest"`
	// The IsMeetngResponsequest element indicates whether incoming messages must be a meeting response in order for the condition or exception to apply.
	IsMeetingResponse *IsMeetingResponse `xml:"IsMeetingResponse"`
	// The IsNDR element indicates whether incoming messages must be non-delivery reports (NDRs) in order for the condition or exception to apply.
	IsNDR *IsNDR `xml:"IsNDR"`
	// The IsPermissionControlled element indicates whether incoming messages must be permission controlled (RMS protected) in order for the condition or exception to apply.
	IsPermissionControlled *IsPermissionControlled `xml:"IsPermissionControlled"`
	// The IsReadReceipt element indicates whether incoming messages must be read receipts in order for the condition or exception to apply.
	IsReadReceipt *IsReadReceipt `xml:"IsReadReceipt"`
	// The IsSigned element indicates whether incoming messages must be signed in order for the condition or exception to apply.
	IsSigned *IsSigned `xml:"IsSigned"`
	// The IsVoicemail element indicates whether incoming messages must be voice mail messages in order for the condition or exception to apply.
	IsVoicemail *IsVoicemail `xml:"IsVoicemail"`
	// The ItemClasses element represents the item classes that must be stamped on incoming messages in order for the condition or exception to apply.
	ItemClasses *ItemClasses `xml:"ItemClasses"`
	// The MessageClassifications element represents the message classifications that must be stamped on incoming messages in order for the condition or exception to apply.
	MessageClassifications *MessageClassifications `xml:"MessageClassifications"`
	// The NotSentToMe element indicates whether the owner of the mailbox must not be in the ToRecipients property of the incoming messages in order for the condition or exception to apply.
	NotSentToMe *NotSentToMe `xml:"NotSentToMe"`
	// The Sensitivity element indicates the sensitivity level of an item.
	Sensitivity *Sensitivity `xml:"Sensitivity"`
	// The SentCcMe element indicates whether the owner of the mailbox has to be in the CcRecipients property of incoming messages in order for the condition or exception to apply.
	SentCcMe *SentCcMe `xml:"SentCcMe"`
	// The SentOnlyToMe element indicates whether the owner of the mailbox has to be the only one in the ToRecipients property of incoming messages in order for the condition or exception to apply.
	SentOnlyToMe *SentOnlyToMe `xml:"SentOnlyToMe"`
	// The SentToAddresses element indicates the e-mail addresses that incoming messages have to have been sent to in order for the condition or exception to apply.
	SentToAddresses *SentToAddresses `xml:"SentToAddresses"`
	// The SentToMe element indicates whether the owner of the mailbox has to be in the ToRecipients property of incoming messages in order for the condition or exception to apply.
	SentToMe *SentToMe `xml:"SentToMe"`
	// The SentToOrCcMe element indicates whether the owner of the mailbox has to be in either a ToRecipients or CcRecipients property of incoming messages in order for the condition or exception to apply.
	SentToOrCcMe *SentToOrCcMe `xml:"SentToOrCcMe"`
	// The WithinDateRange element specifies the date range within which incoming messages have to have been received in order for the condition or exception to apply.
	WithinDateRange *WithinDateRange `xml:"WithinDateRange"`
	// The WithinSizeRange element specifies the minimum and maximum sizes that incoming messages must be in order for the condition or exception to apply.
	WithinSizeRange *WithinSizeRange `xml:"WithinSizeRange"`
}

func (E *Exceptions) SetForMarshal() {
	E.XMLName.Local = "t:Exceptions"
}

func (E *Exceptions) GetSchema() *Schema {
	return &SchemaTypes
}
