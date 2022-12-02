package elements

// The RightsManagementLicenseData element specifies information about the rights management license for an item.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/rightsmanagementlicensedata
type RightsManagementLicenseData struct {
	// The ContentExpiryDate element specifies the expiration date of the content.
	ContentExpiryDate *ContentExpiryDate `xml:"t:ContentExpiryDate"`
	// The ContentOwner element specifies the name of the content owner.
	ContentOwner *ContentOwner `xml:"t:ContentOwner"`
	// The EditAllowed element specifies whether Information Rights Management can be edited.
	EditAllowed *EditAllowed `xml:"t:EditAllowed"`
	// The ExportAllowed element specifies whether exporting is enabled.
	ExportAllowed *ExportAllowed `xml:"t:ExportAllowed"`
	// The ExtractAllowed element specifies whether entity extraction is enabled.
	ExtractAllowed *ExtractAllowed `xml:"t:ExtractAllowed"`
	// The ForwardAllowed element specifies whether forwarding emails is enabled.
	ForwardAllowed *ForwardAllowed `xml:"t:ForwardAllowed"`
	// The IsOwner element specifies whether the specified email user is the owner.
	IsOwner *IsOwner `xml:"t:IsOwner"`
	// The ModifyRecipientsAllowed element specifies whether modification of the recipients is enabled.
	ModifyRecipientsAllowed *ModifyRecipientsAllowed `xml:"t:ModifyRecipientsAllowed"`
	// The PrintAllowed element specifies whether printing is enabled.
	PrintAllowed *PrintAllowed `xml:"t:PrintAllowed"`
	// The ProgrammaticAccessAllowed element specifies whether programmatic access is enabled for rights managed data.
	ProgrammaticAccessAllowed *ProgrammaticAccessAllowed `xml:"t:ProgrammaticAccessAllowed"`
	// The RMSTemplateId element specifies the identifier of the Rights Management template.
	RMSTemplateId *RMSTemplateId `xml:"t:RMSTemplateId"`
	// The ReplyAllAllowed element specifies whether a reply all is allowed for rights managed data.
	ReplyAllAllowed *ReplyAllAllowed `xml:"t:ReplyAllAllowed"`
	// The ReplyAllowed element specifies whether a reply is allowed for rights managed data.
	ReplyAllowed *ReplyAllowed `xml:"t:ReplyAllowed"`
	// The RightsManagedMessageDecryptionStatus element specifies the rights management decryption status of an item.
	RightsManagedMessageDecryptionStatus *RightsManagedMessageDecryptionStatus `xml:"t:RightsManagedMessageDecryptionStatus"`
	// The TemplateDescription element specifies the description of the Rights Management template.
	TemplateDescription *TemplateDescription `xml:"t:TemplateDescription"`
	// The TemplateName element specifies the name of the Rights Management template.
	TemplateName *TemplateName `xml:"t:TemplateName"`
}
