package elements

// The Contact element represents a contact item in the Exchange store.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/contact
import "encoding/xml"

type Contact struct {
	XMLName xml.Name

	// The Alias element contains the email alias of a contact.
	Alias *Alias `xml:"Alias"`
	// The AssitantName element represents an assistant to a contact.
	AssistantName *AssistantName `xml:"AssistantName"`
	// The Attachments element contains the items or files that are attached to an item in the Exchange store.
	Attachments *Attachments `xml:"Attachments"`
	// The Birthday element represents the birth date of a contact.
	Birthday *Birthday `xml:"Birthday"`
	// The Body element specifies the body of an item.
	Body *Body `xml:"Body"`
	// The BusinessHomePage element represents the Home page (Web address) for the contact.
	BusinessHomePage *BusinessHomePage `xml:"BusinessHomePage"`
	// The Categories element contains a collection of strings that identify the categories to which an item in the mailbox belongs.
	Categories *Categories `xml:"Categories"`
	// The Children element contains the names of a contact's children.
	Children *Children `xml:"Children"`
	// The Companies element represents the collection of companies that are associated with a contact or task.
	Companies *Companies `xml:"Companies"`
	// The CompanyName element represents the company name that is associated with a contact.
	CompanyName *CompanyName `xml:"CompanyName"`
	// The CompleteName element represents the complete name of a contact.
	CompleteName *CompleteName `xml:"CompleteName"`
	// The ContactSource element describes whether the contact is located in the Exchange store or Active Directory Domain Services (AD DS).
	ContactSource *ContactSource `xml:"ContactSource"`
	// The ConversationId element contains the identifier of an item or conversation.
	ConversationId *ConversationId `xml:"ConversationId"`
	// The Culture element represents the culture for a given item in a mailbox.
	Culture *Culture `xml:"Culture"`
	// The DateTimeCreated element represents the date and time that an item in the mailbox was created.
	DateTimeCreated *DateTimeCreated `xml:"DateTimeCreated"`
	// The DateTimeReceived element represents the date and time that an item in a mailbox was received.
	DateTimeReceived *DateTimeReceived `xml:"DateTimeReceived"`
	// The DateTimeSent element represents the date and time at which an item in a mailbox was sent.
	DateTimeSent *DateTimeSent `xml:"DateTimeSent"`
	// The Department element represents the contact's department at work.
	Department *Department `xml:"Department"`
	// The DirectReports element contains SMTP information that identifies the direct reports of a contact.
	DirectReports *DirectReports `xml:"DirectReports"`
	// The DirectoryId element contains the directory ID of a contact.
	DirectoryId *DirectoryId `xml:"DirectoryId"`
	// The DisplayCc element represents the display string that is used for the contents of the Cc box. This is the concatenated string of all Cc recipient display names.
	DisplayCc *DisplayCc `xml:"DisplayCc"`
	// The DisplayName element defines the display name of a folder, contact, distribution list, delegate user, location, or rule.
	DisplayName *DisplayNamestring `xml:"DisplayName"`
	// The DisplayTo element represents the display string that is used for the contents of the To box. This is the concatenated string of all To recipient display names.
	DisplayTo *DisplayTo `xml:"DisplayTo"`
	// The EffectiveRights element contains the client's rights based on the permission settings for the item or folder. This element is read-only.
	EffectiveRights *EffectiveRights `xml:"EffectiveRights"`
	// The EmailAddresses element represents a collection of e-mail addresses for a contact.
	EmailAddresses *EmailAddresses `xml:"EmailAddresses"`
	// The ExtendedProperty element identifies extended MAPI properties on folders and items.
	ExtendedProperty *ExtendedProperty `xml:"ExtendedProperty"`
	// The FileAs element represents how a contact or distribution list is filed in the Contacts folder.
	FileAs *FileAs `xml:"FileAs"`
	// The FileAsMapping element defines how to construct what is displayed for a contact.
	FileAsMapping *FileAsMapping `xml:"FileAsMapping"`
	// The Generation element represents a generational abbreviation that follows the full name of a contact.
	Generation *Generation `xml:"Generation"`
	// The GivenName element contains a contact's given name.
	GivenName *GivenName `xml:"GivenName"`
	// The HasAttachments element represents a property that is set to true if an item has at least one visible attachment or if a conversation contains at least one item that has an attachment. This property is read-only.
	HasAttachments *HasAttachments `xml:"HasAttachments"`
	// The HasPicture element indicates whether the contact item has a file attachment that represents the contact's picture.
	HasPicture *HasPicture `xml:"HasPicture"`
	// The ImAddresses element represents a collection of instant messaging addresses for a contact.
	ImAddresses *ImAddresses `xml:"ImAddresses"`
	// The Importance element describes the importance of an item or the aggregated importance of all items in a conversation in the current folder.
	Importance *Importance `xml:"Importance"`
	// The InReplyTo element represents the identifier of the item to which this item is a reply.
	InReplyTo *InReplyTo `xml:"InReplyTo"`
	// The Initials element represents the initials of a contact.
	Initials *Initials `xml:"Initials"`
	// The InternetMessageHeaders element contains a collection of some of the Internet message headers that are contained in an item in a mailbox. To get the entire collection of Internet message headers, use the PR_TRANSPORT_MESSAGE_HEADERS property. For more information about EWS and Internet message headers, seeGetting Internet message headersin EWS, MIME, and the missing Internet message headers.
	InternetMessageHeaders *InternetMessageHeaders `xml:"InternetMessageHeaders"`
	// The IsAssociated element indicates whether the item is associated with a folder.
	IsAssociated *IsAssociated `xml:"IsAssociated"`
	// The IsDraft element indicates whether an item has not yet been sent.
	IsDraft *IsDraft `xml:"IsDraft"`
	// The IsFromMe element indicates whether a user sent an item to him or herself.
	IsFromMe *IsFromMe `xml:"IsFromMe"`
	// The IsResend element indicates whether the item had previously been sent.
	IsResend *IsResend `xml:"IsResend"`
	// The IsSubmitted element indicates whether an item has been submitted to the Outbox default folder.
	IsSubmitted *IsSubmitted `xml:"IsSubmitted"`
	// The IsUnmodified element indicates whether the item has been modified.
	IsUnmodified *IsUnmodified `xml:"IsUnmodified"`
	// The ItemClass element represents the message class of an item.
	ItemClass *ItemClass `xml:"ItemClass"`
	// The ItemId element contains the unique identifier and change key of an item in the Exchange store.
	ItemId *ItemId `xml:"ItemId"`
	// The JobTitle element represents the job title of a contact.
	JobTitle *JobTitle `xml:"JobTitle"`
	// The LastModifiedName element contains the display name of the last user to modify an item. This element is read-only. This element was introduced in Microsoft Exchange Server 2007 Service Pack 1 (SP1).
	LastModifiedName *LastModifiedName `xml:"LastModifiedName"`
	// The LastModifiedTime element indicates when an item was last modified. This element is read-only.
	LastModifiedTime *LastModifiedTime `xml:"LastModifiedTime"`
	// The MSExchangeCertificate element contains a value that encodes the Microsoft Exchange certificate of a contact.
	MSExchangeCertificate *MSExchangeCertificate `xml:"MSExchangeCertificate"`
	// The Manager element represents a contact's manager.
	Manager *Manager `xml:"Manager"`
	// The ManagerMailbox element contains SMTP information that identifies the mailbox of the contact's manager.
	ManagerMailbox *ManagerMailbox `xml:"ManagerMailbox"`
	// The MiddleName element represents the middle name of a contact.
	MiddleName *MiddleName `xml:"MiddleName"`
	// The Mileage element represents mileage for a task or contact item.
	Mileage *Mileage `xml:"Mileage"`
	// The MimeContent element contains the ASCII MIME stream of an object that is represented in base64Binary format and supports [RFC2045].
	MimeContent *MimeContent `xml:"MimeContent"`
	// The Nickname element represents the nickname of a contact.
	Nickname *Nickname `xml:"Nickname"`
	// The Notes element contains supplementary contact information.
	Notes *NotesContact `xml:"Notes"`
	// The OfficeLocation element represents the office location of a contact.
	OfficeLocation *OfficeLocation `xml:"OfficeLocation"`
	// The ParentFolderId element represents the identifier of the parent folder that contains the item or folder.
	ParentFolderId *ParentFolderId `xml:"ParentFolderId"`
	// The PhoneNumbers element represents a collection of telephone numbers for a contact.
	PhoneNumbers *PhoneNumbers `xml:"PhoneNumbers"`
	// The PhoneticFirstName element contains the first name of a contact, spelled phonetically.
	PhoneticFirstName *PhoneticFirstName `xml:"PhoneticFirstName"`
	// The PhoneticFullName element contains the full name of a contact, including the first and last name, spelled phonetically.
	PhoneticFullName *PhoneticFullName `xml:"PhoneticFullName"`
	// The PhoneticLastName element contains the last name of a contact, spelled phonetically.
	PhoneticLastName *PhoneticLastName `xml:"PhoneticLastName"`
	// The Photo element contains a value that encodes the photo of a contact.
	Photo *Photo `xml:"Photo"`
	// The PhysicalAddresses element contains a collection of physical addresses that are associated with a contact.
	PhysicalAddresses *PhysicalAddresses `xml:"PhysicalAddresses"`
	// The PostalAddressIndex element represents the display types for physical addresses.
	PostalAddressIndex *PostalAddressIndex `xml:"PostalAddressIndex"`
	// The Profession element represents the profession of a contact.
	Profession *Profession `xml:"Profession"`
	// The ReminderDueBy element represents the date and time when the event occurs. This is used by the ReminderMinutesBeforeStart element to determine when the reminder is displayed.
	ReminderDueBy *ReminderDueBy `xml:"ReminderDueBy"`
	// The ReminderIsSet element indicates whether a reminder has been set for an item in the Exchange store.
	ReminderIsSet *ReminderIsSet `xml:"ReminderIsSet"`
	// The ReminderMinutesBeforeStart element represents the number of minutes before an event occurs when a reminder is displayed.
	ReminderMinutesBeforeStart *ReminderMinutesBeforeStart `xml:"ReminderMinutesBeforeStart"`
	// The ResponseObjects element contains a collection of all the response objects that are associated with an item in the Exchange store.
	ResponseObjects *ResponseObjects `xml:"ResponseObjects"`
	// The Sensitivity element indicates the sensitivity level of an item.
	Sensitivity *Sensitivity `xml:"Sensitivity"`
	// The Size element represents the size in bytes of an item or all the items in a conversation in the current folder. This property is read-only.
	Size *Size `xml:"Size"`
	// The SpouseName element represents the name of a contact's spouse or partner.
	SpouseName *SpouseName `xml:"SpouseName"`
	// The Subject element represents the subject property of Exchange store items. The subject is limited to 255 characters.
	Subject *Subject `xml:"Subject"`
	// The Surname element represents the surname of a contact.
	Surname *Surname `xml:"Surname"`
	// The UniqueBody element represents an HTML fragment or plain text which represents the unique body of this conversation.
	UniqueBody *UniqueBody `xml:"UniqueBody"`
	// The UserSMIMECertificate element contains a value that encodes a contact's SMIME certificate.
	UserSMIMECertificate *UserSMIMECertificate `xml:"UserSMIMECertificate"`
	// The WebClientEditFormQueryString element represents a URL to concatenate to the Outlook Web App endpoint to edit an item in Outlook Web App.
	WebClientEditFormQueryString *WebClientEditFormQueryString `xml:"WebClientEditFormQueryString"`
	// The WebClientReadFormQueryString element represents a URL to concatenate to the Outlook Web App endpoint to read an item in Outlook Web App.
	WebClientReadFormQueryString *WebClientReadFormQueryString `xml:"WebClientReadFormQueryString"`
	// The WeddingAnniversary element contains the wedding anniversary of a contact.
	WeddingAnniversary *WeddingAnniversary `xml:"WeddingAnniversary"`
}

func (C *Contact) SetForMarshal() {
	C.XMLName.Local = "t:Contact"
}

func (C *Contact) GetSchema() *Schema {
	return &SchemaTypes
}
