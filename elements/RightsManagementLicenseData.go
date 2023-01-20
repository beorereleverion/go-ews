package elements

// The RightsManagementLicenseData element specifies information about the rights management license for an item.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/rightsmanagementlicensedata
import "encoding/xml"

type RightsManagementLicenseData struct {
	XMLName xml.Name

	// The ContentExpiryDate element specifies the expiration date of the content.
	ContentExpiryDate *ContentExpiryDate `xml:"ContentExpiryDate"`
	// The ContentOwner element specifies the name of the content owner.
	ContentOwner *ContentOwner `xml:"ContentOwner"`
	// The EditAllowed element specifies whether Information Rights Management can be edited.
	EditAllowed *EditAllowed `xml:"EditAllowed"`
	// The ExportAllowed element specifies whether exporting is enabled.
	ExportAllowed *ExportAllowed `xml:"ExportAllowed"`
	// The ExtractAllowed element specifies whether entity extraction is enabled.
	ExtractAllowed *ExtractAllowed `xml:"ExtractAllowed"`
	// The ForwardAllowed element specifies whether forwarding emails is enabled.
	ForwardAllowed *ForwardAllowed `xml:"ForwardAllowed"`
	// The IsOwner element specifies whether the specified email user is the owner.
	IsOwner *IsOwner `xml:"IsOwner"`
	// The ModifyRecipientsAllowed element specifies whether modification of the recipients is enabled.
	ModifyRecipientsAllowed *ModifyRecipientsAllowed `xml:"ModifyRecipientsAllowed"`
	// The PrintAllowed element specifies whether printing is enabled.
	PrintAllowed *PrintAllowed `xml:"PrintAllowed"`
	// The ProgrammaticAccessAllowed element specifies whether programmatic access is enabled for rights managed data.
	ProgrammaticAccessAllowed *ProgrammaticAccessAllowed `xml:"ProgrammaticAccessAllowed"`
	// The RMSTemplateId element specifies the identifier of the Rights Management template.
	RMSTemplateId *RMSTemplateId `xml:"RMSTemplateId"`
	// The ReplyAllAllowed element specifies whether a reply all is allowed for rights managed data.
	ReplyAllAllowed *ReplyAllAllowed `xml:"ReplyAllAllowed"`
	// The ReplyAllowed element specifies whether a reply is allowed for rights managed data.
	ReplyAllowed *ReplyAllowed `xml:"ReplyAllowed"`
	// The RightsManagedMessageDecryptionStatus element specifies the rights management decryption status of an item.
	RightsManagedMessageDecryptionStatus *RightsManagedMessageDecryptionStatus `xml:"RightsManagedMessageDecryptionStatus"`
	// The TemplateDescription element specifies the description of the Rights Management template.
	TemplateDescription *TemplateDescription `xml:"TemplateDescription"`
	// The TemplateName element specifies the name of the Rights Management template.
	TemplateName *TemplateName `xml:"TemplateName"`
}

func (R *RightsManagementLicenseData) SetForMarshal() {
	R.XMLName.Local = "t:RightsManagementLicenseData"
}

func (R *RightsManagementLicenseData) GetSchema() *Schema {
	return &SchemaTypes
}
