package elements

// The Actions element contains a list of actions associated with Inbox rules.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/actions-ruleactionstype
type ActionsRuleActionsType struct {
	// The AssignCategories element represents the categories that are stamped on e-mail messages.
	AssignCategories *AssignCategories `xml:"m:AssignCategories"`
	// The CopyToFolder element specifies the identifier of the folder that email items can be copied to.
	CopyToFolder *CopyToFolder `xml:"m:CopyToFolder"`
	// The Delete element indicates whether a client can delete a folder or item.
	Delete *Delete `xml:"t:Delete"`
	// The ForwardAsAttachmentToRecipients element indicates the e-mail addresses to which messages are to be forwarded as attachments.
	ForwardAsAttachmentToRecipients *ForwardAsAttachmentToRecipients `xml:"m:ForwardAsAttachmentToRecipients"`
	// The ForwardToRecipients element indicates the e-mail addresses to which messages are to be forwarded.
	ForwardToRecipients *ForwardToRecipients `xml:"m:ForwardToRecipients"`
	// The MarkAsRead element indicates whether messages are to be marked as read.
	MarkAsRead *MarkAsRead `xml:"m:MarkAsRead"`
	// The MarkImportance element specifies the importance that is to be stamped on messages.
	MarkImportance *MarkImportance `xml:"m:MarkImportance"`
	// The MoveToFolder element specifies the identifier of the folder to which email items can be moved.
	MoveToFolder *MoveToFolder `xml:"m:MoveToFolder"`
	// The PermanentDelete element indicates whether messages are to be permanently deleted and not saved to the Deleted Items folder.
	PermanentDelete *PermanentDelete `xml:"m:PermanentDelete"`
	// The RedirectToRecipients element indicates the e-mail addresses to which messages are to be redirected.
	RedirectToRecipients *RedirectToRecipients `xml:"m:RedirectToRecipients"`
	// The SendSMSAlertToRecipients element indicates the mobile phone numbers to which a Short Message Service (SMS) alert is to be sent.
	SendSMSAlertToRecipients *SendSMSAlertToRecipients `xml:"m:SendSMSAlertToRecipients"`
	// The ServerReplyWithMessage element indicates the ID of the template message that is to be sent as a reply to incoming messages.
	ServerReplyWithMessage *ServerReplyWithMessage `xml:"m:ServerReplyWithMessage"`
	// The StopProcessingRules element indicates whether subsequent rules are to be evaluated.
	StopProcessingRules *StopProcessingRules `xml:"m:StopProcessingRules"`
}
