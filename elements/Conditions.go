package elements

// The Conditions element identifies the conditions that, when fulfilled, will trigger the rule actions for a rule.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/conditions
type Conditions struct {
	// The Categories element contains a collection of strings that identify the categories to which an item in the mailbox belongs.
	Categories *Categories `xml:"t:Categories"`
	// The ContainsBodyStrings element indicates the strings that must appear in the body of incoming messages in order for the condition or exception to apply.
	ContainsBodyStrings *ContainsBodyStrings `xml:"m:ContainsBodyStrings"`
	// The ContainsHeaderStrings element indicates the strings that must appear in the headers of incoming messages in order for the condition or exception to apply.
	ContainsHeaderStrings *ContainsHeaderStrings `xml:"m:ContainsHeaderStrings"`
	// The ContainsRecipientStrings element indicates the strings that must appear in either the ToRecipients or CcRecipients properties of incoming messages in order for the condition or exception to apply.
	ContainsRecipientStrings *ContainsRecipientStrings `xml:"m:ContainsRecipientStrings"`
	// The ContainsSenderStrings element indicates the strings that must appear in the From property of incoming messages in order for the condition or exception to apply.
	ContainsSenderStrings *ContainsSenderStrings `xml:"m:ContainsSenderStrings"`
	// The ContainsSubjectOrBodyStrings element indicates the strings that must appear in either the body or the subject of incoming messages in order for the condition or exception to apply.
	ContainsSubjectOrBodyStrings *ContainsSubjectOrBodyStrings `xml:"m:ContainsSubjectOrBodyStrings"`
	// The ContainsSubjectStrings element indicates the strings that must appear in the subject of incoming messages in order for the condition or exception to apply.
	ContainsSubjectStrings *ContainsSubjectStrings `xml:"m:ContainsSubjectStrings"`
	// The FlaggedForAction element specifies the flag for action value that must appear on incoming messages in order for the condition or exception to apply.
	FlaggedForAction *FlaggedForAction `xml:"m:FlaggedForAction"`
	// The FromAddresses element indicates the e-mail addresses from which incoming messages must be sent in order for the condition or exception to apply.
	FromAddresses *FromAddresses `xml:"m:FromAddresses"`
	// The FromConnectedAccounts element represents the e-mail account names from which incoming messages have to have been aggregated in order for the condition or exception to apply.
	FromConnectedAccounts *FromConnectedAccounts `xml:"m:FromConnectedAccounts"`
	// The HasAttachments element represents a property that is set to true if an item has at least one visible attachment or if a conversation contains at least one item that has an attachment. This property is read-only.
	HasAttachments *HasAttachments `xml:"t:HasAttachments"`
	// The Importance element describes the importance of an item or the aggregated importance of all items in a conversation in the current folder.
	Importance *Importance `xml:"t:Importance"`
	// The IsApprovalRequest element indicates whether incoming messages must be approval requests in order for the condition or exception to apply.
	IsApprovalRequest *IsApprovalRequest `xml:"m:IsApprovalRequest"`
	// The IsAutomaticForward element indicates whether incoming messages must be automatic forwards in order for the condition or exception to apply.
	IsAutomaticForward *IsAutomaticForward `xml:"m:IsAutomaticForward"`
	// The IsAutomaticReply element indicates whether incoming messages must be automatic replies in order for the condition or exception to apply.
	IsAutomaticReply *IsAutomaticReply `xml:"m:IsAutomaticReply"`
	// The IsEncrypted element indicates whether incoming messages must be S/MIME encrypted in order for the condition or exception to apply.
	IsEncrypted *IsEncrypted `xml:"m:IsEncrypted"`
	// The IsMeetngRequest element indicates whether incoming messages must be a meeting request in order for the condition or exception to apply.
	IsMeetingRequest *IsMeetingRequest `xml:"m:IsMeetingRequest"`
	// The IsMeetngResponsequest element indicates whether incoming messages must be a meeting response in order for the condition or exception to apply.
	IsMeetingResponse *IsMeetingResponse `xml:"m:IsMeetingResponse"`
	// The IsNDR element indicates whether incoming messages must be non-delivery reports (NDRs) in order for the condition or exception to apply.
	IsNDR *IsNDR `xml:"m:IsNDR"`
	// The IsPermissionControlled element indicates whether incoming messages must be permission controlled (RMS protected) in order for the condition or exception to apply.
	IsPermissionControlled *IsPermissionControlled `xml:"m:IsPermissionControlled"`
	// The IsReadReceipt element indicates whether incoming messages must be read receipts in order for the condition or exception to apply.
	IsReadReceipt *IsReadReceipt `xml:"m:IsReadReceipt"`
	// The IsSigned element indicates whether incoming messages must be signed in order for the condition or exception to apply.
	IsSigned *IsSigned `xml:"m:IsSigned"`
	// The IsVoicemail element indicates whether incoming messages must be voice mail messages in order for the condition or exception to apply.
	IsVoicemail *IsVoicemail `xml:"m:IsVoicemail"`
	// The ItemClasses element represents the item classes that must be stamped on incoming messages in order for the condition or exception to apply.
	ItemClasses *ItemClasses `xml:"t:ItemClasses"`
	// The MessageClassifications element represents the message classifications that must be stamped on incoming messages in order for the condition or exception to apply.
	MessageClassifications *MessageClassifications `xml:"m:MessageClassifications"`
	// The NotSentToMe element indicates whether the owner of the mailbox must not be in the ToRecipients property of the incoming messages in order for the condition or exception to apply.
	NotSentToMe *NotSentToMe `xml:"m:NotSentToMe"`
	// The Sensitivity element indicates the sensitivity level of an item.
	Sensitivity *Sensitivity `xml:"t:Sensitivity"`
	// The SentCcMe element indicates whether the owner of the mailbox has to be in the CcRecipients property of incoming messages in order for the condition or exception to apply.
	SentCcMe *SentCcMe `xml:"m:SentCcMe"`
	// The SentOnlyToMe element indicates whether the owner of the mailbox has to be the only one in the ToRecipients property of incoming messages in order for the condition or exception to apply.
	SentOnlyToMe *SentOnlyToMe `xml:"m:SentOnlyToMe"`
	// The SentToAddresses element indicates the e-mail addresses that incoming messages have to have been sent to in order for the condition or exception to apply.
	SentToAddresses *SentToAddresses `xml:"m:SentToAddresses"`
	// The SentToMe element indicates whether the owner of the mailbox has to be in the ToRecipients property of incoming messages in order for the condition or exception to apply.
	SentToMe *SentToMe `xml:"m:SentToMe"`
	// The SentToOrCcMe element indicates whether the owner of the mailbox has to be in either a ToRecipients or CcRecipients property of incoming messages in order for the condition or exception to apply.
	SentToOrCcMe *SentToOrCcMe `xml:"m:SentToOrCcMe"`
	// The WithinDateRange element specifies the date range within which incoming messages have to have been received in order for the condition or exception to apply.
	WithinDateRange *WithinDateRange `xml:"m:WithinDateRange"`
	// The WithinSizeRange element specifies the minimum and maximum sizes that incoming messages must be in order for the condition or exception to apply.
	WithinSizeRange *WithinSizeRange `xml:"m:WithinSizeRange"`
}
