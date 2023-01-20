package elements

// The Actions element contains a list of actions associated with Inbox rules.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/actions-ruleactionstype
import "encoding/xml"

type ActionsRuleActionsType struct {
	XMLName xml.Name

	// The AssignCategories element represents the categories that are stamped on e-mail messages.
	AssignCategories *AssignCategories `xml:"AssignCategories"`
	// The CopyToFolder element specifies the identifier of the folder that email items can be copied to.
	CopyToFolder *CopyToFolder `xml:"CopyToFolder"`
	// The Delete element indicates whether a client can delete a folder or item.
	Delete *Delete `xml:"Delete"`
	// The ForwardAsAttachmentToRecipients element indicates the e-mail addresses to which messages are to be forwarded as attachments.
	ForwardAsAttachmentToRecipients *ForwardAsAttachmentToRecipients `xml:"ForwardAsAttachmentToRecipients"`
	// The ForwardToRecipients element indicates the e-mail addresses to which messages are to be forwarded.
	ForwardToRecipients *ForwardToRecipients `xml:"ForwardToRecipients"`
	// The MarkAsRead element indicates whether messages are to be marked as read.
	MarkAsRead *MarkAsRead `xml:"MarkAsRead"`
	// The MarkImportance element specifies the importance that is to be stamped on messages.
	MarkImportance *MarkImportance `xml:"MarkImportance"`
	// The MoveToFolder element specifies the identifier of the folder to which email items can be moved.
	MoveToFolder *MoveToFolder `xml:"MoveToFolder"`
	// The PermanentDelete element indicates whether messages are to be permanently deleted and not saved to the Deleted Items folder.
	PermanentDelete *PermanentDelete `xml:"PermanentDelete"`
	// The RedirectToRecipients element indicates the e-mail addresses to which messages are to be redirected.
	RedirectToRecipients *RedirectToRecipients `xml:"RedirectToRecipients"`
	// The SendSMSAlertToRecipients element indicates the mobile phone numbers to which a Short Message Service (SMS) alert is to be sent.
	SendSMSAlertToRecipients *SendSMSAlertToRecipients `xml:"SendSMSAlertToRecipients"`
	// The ServerReplyWithMessage element indicates the ID of the template message that is to be sent as a reply to incoming messages.
	ServerReplyWithMessage *ServerReplyWithMessage `xml:"ServerReplyWithMessage"`
	// The StopProcessingRules element indicates whether subsequent rules are to be evaluated.
	StopProcessingRules *StopProcessingRules `xml:"StopProcessingRules"`
}

func (A *ActionsRuleActionsType) SetForMarshal() {
	A.XMLName.Local = "t:Actions"
}

func (A *ActionsRuleActionsType) GetSchema() *Schema {
	return &SchemaTypes
}
