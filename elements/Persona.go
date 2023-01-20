package elements

// The Persona element specifies a set of persona data returned by a GetPersona request.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/persona
import "encoding/xml"

type Persona struct {
	XMLName xml.Name

	// The AssistantNames element specifies an array of assistant names and the identifiers of their source attributions for the associated persona.
	AssistantNames *AssistantNames `xml:"AssistantNames"`
	// The AssistantPhoneNumbers element specifies an array of assistant phone numbers and the identifiers of their source attributions for the associated persona.
	AssistantPhoneNumbers *AssistantPhoneNumbers `xml:"AssistantPhoneNumbers"`
	// The Attributions element specifies an array of attribution information for one or more of the contacts or Active Directory recipients aggregated into the associated persona.
	Attributions *AttributionsArrayOfPersonaAttributionsType `xml:"Attributions"`
	// The Birthdays element specifies an array of birthdays, stored as strings, and the identifiers of their source attributions for the associated persona.
	Birthdays *Birthdays `xml:"Birthdays"`
	// The Bodies element specifies an array of BodyContentAttributedValue elements.
	Bodies *Bodies `xml:"Bodies"`
	// The BusinessAddresses element specifies an array of business addresses and the identifiers of their source attributions for the associated persona.
	BusinessAddresses *BusinessAddresses `xml:"BusinessAddresses"`
	// The BusinessHomePages element specifies an array of business home pages and the identifiers of their source attributions for the associated persona.
	BusinessHomePages *BusinessHomePages `xml:"BusinessHomePages"`
	// The BusinessPhoneNumbers element specifies an array of business phone numbers and the identifiers of their source attributions for the associated persona.
	BusinessPhoneNumbers *BusinessPhoneNumbers `xml:"BusinessPhoneNumbers"`
	// The BusinessPhoneNumbers2 element specifies an array of BusinessPhoneNumber2 elements and the identifiers of their source attributions for the associated persona.
	BusinessPhoneNumbers2 *BusinessPhoneNumbers2 `xml:"BusinessPhoneNumbers2"`
	// The CallbackPhones element specifies an array of call-back phone numbers and the identifiers of their source attributions for the associated persona.
	CallbackPhones *CallbackPhones `xml:"CallbackPhones"`
	// The CarPhone element specifies an array of car phone numbers and the identifiers of their source attributions for the associated persona.
	CarPhones *CarPhones `xml:"CarPhones"`
	// The Children element specifies an array of child names and identifiers of their source attributions for the associated persona.
	Children *ChildrenArrayOfStringArrayAttributedValuesType `xml:"Children"`
	// The CompanyName element represents the company name that is associated with a contact.
	CompanyName *CompanyName `xml:"CompanyName"`
	// The CompanyNameSortKey element contains the sort key for a company name.
	CompanyNameSortKey *CompanyNameSortKey `xml:"CompanyNameSortKey"`
	// The CompanyNames element contains an array of company names and the identifiers of their source attributions for the associated persona.
	CompanyNames *CompanyNames `xml:"CompanyNames"`
	// The CreationTime element specifies when the persona was created.
	CreationTime *CreationTime `xml:"CreationTime"`
	// The Department element represents the contact's department at work.
	Department *Department `xml:"Department"`
	// The Departments element specifies an array of department names and the identifiers of their source attributions for the associated persona.
	Departments *Departments `xml:"Departments"`
	// The DisplayName element defines the display name of a folder, contact, distribution list, delegate user, location, or rule.
	DisplayName *DisplayNamestring `xml:"DisplayName"`
	// The DisplayNameFirstLast element specifies the display name of the associated persona in the format,First Name,Last Name.
	DisplayNameFirstLast *DisplayNameFirstLast `xml:"DisplayNameFirstLast"`
	// The DisplayNameFirstLastHeader element specifies the header for the display name, first name first.
	DisplayNameFirstLastHeader *DisplayNameFirstLastHeader `xml:"DisplayNameFirstLastHeader"`
	// The DisplayNameFirstLastSortKey element contains the sort key for a display name in first name, last name order.
	DisplayNameFirstLastSortKey *DisplayNameFirstLastSortKey `xml:"DisplayNameFirstLastSortKey"`
	// The DisplayNameLastFirst element specifies the display name of the associated persona in the format,Last Name,First Name.
	DisplayNameLastFirst *DisplayNameLastFirst `xml:"DisplayNameLastFirst"`
	// The DisplayNameLastFirstHeader element specifies the header for the display name, last name first.
	DisplayNameLastFirstHeader *DisplayNameLastFirstHeader `xml:"DisplayNameLastFirstHeader"`
	// The DisplayNameLastFirstSortKey element contains the sort key for a display name in last name, first name order.
	DisplayNameLastFirstSortKey *DisplayNameLastFirstSortKey `xml:"DisplayNameLastFirstSortKey"`
	// The DisplayNamePrefix element specifies the prefix of the display name of the associated persona.
	DisplayNamePrefix *DisplayNamePrefix `xml:"DisplayNamePrefix"`
	// The DisplayNamePrefixes element specifies an array of display name prefixes and the identifiers of their source attributions for the associated persona.
	DisplayNamePrefixes *DisplayNamePrefixes `xml:"DisplayNamePrefixes"`
	// The DisplayNames element specifies an array of display names and the identifiers of their source attributions for the associated persona.
	DisplayNames *DisplayNames `xml:"DisplayNames"`
	// The EmailAddress element specifies the fully resolved SMTP address for the site mailbox or the associated persona.
	EmailAddress *EmailAddressEmailAddressType `xml:"EmailAddress"`
	// The EmailAddresses element specifies an array of all email addresses of the associated persona.
	EmailAddresses *EmailAddressesArrayOfEmailAddressesType `xml:"EmailAddresses"`
	// The Emails1 element specifies an array of EmailAddressAttributedValue values and the identifiers of their source attributions for the associated persona.
	Emails1 *Emails1 `xml:"Emails1"`
	// The Emails2 element contains an array of EmailAddressAttributedValue values and the identifiers of their source attributions for the associated persona.
	Emails2 *Emails2 `xml:"Emails2"`
	// The Emails3 element specifies an array of EmailAddressAttributedValue values and the identifiers of their source attributions for the associated persona.
	Emails3 *Emails3 `xml:"Emails3"`
	// The ExtendedProperties element contains the extended properties used for a persona.
	ExtendedProperties *ExtendedPropertiesArrayOfExtendedPropertyAttributedValueType `xml:"ExtendedProperties"`
	// The FileAs element represents how a contact or distribution list is filed in the Contacts folder.
	FileAs *FileAs `xml:"FileAs"`
	// The FileAsHeader specifies the header for the File As option.
	FileAsHeader *FileAsHeader `xml:"FileAsHeader"`
	// The FileAsId element specifies the FileAs identifier.
	FileAsId *FileAsId `xml:"FileAsId"`
	// The FileAsIds element specifies an array of StringAttributedValue elements and the identifiers of their source attributions for the associated persona.
	FileAsIds *FileAsIds `xml:"FileAsIds"`
	// The FileAses element specifies an array of StringAttributedValue elements and the identifiers of their source attributions for the associated persona.
	FileAses *FileAses `xml:"FileAses"`
	// The FolderIds element contains a list of folder identifiers.
	FolderIds *FolderIdsArrayOfFolderIdType `xml:"FolderIds"`
	// The Generation element represents a generational abbreviation that follows the full name of a contact.
	Generation *Generation `xml:"Generation"`
	// The Generations element specifies an array of generation values and the identifiers of their source attributions for the associated persona.
	Generations *Generations `xml:"Generations"`
	// The GivenName element contains a contact's given name.
	GivenName *GivenName `xml:"GivenName"`
	// The GivenNames element specifies an array of given name values and the identifiers of their source attributions for the associated persona.
	GivenNames *GivenNames `xml:"GivenNames"`
	// The Hobbies element specifies an array of hobbies and the identifiers of their source attributions for the associated persona.
	Hobbies *Hobbies `xml:"Hobbies"`
	// The HomeAddresses element specifies an array of home addresses and the identifiers of their source attributions for the associated persona.
	HomeAddresses *HomeAddresses `xml:"HomeAddresses"`
	// The HomeCity element specifies the city of the home address of the associated persona.
	HomeCity *HomeCity `xml:"HomeCity"`
	// The HomeCitySortKey element represents the sort key for the home city.
	HomeCitySortKey *HomeCitySortKey `xml:"HomeCitySortKey"`
	// The HomeFaxes element specifies an array of home fax numbers and the identifiers of their source attributions for the associated persona.
	HomeFaxes *HomeFaxes `xml:"HomeFaxes"`
	// The HomePhones element specifies an array of home phone numbers and the identifiers of their source attributions for the associated persona.
	HomePhones *HomePhones `xml:"HomePhones"`
	// The HomePhones2 element specifies an array of HomePhone2 values and the identifiers of their source attributions for the associated persona.
	HomePhones2 *HomePhones2 `xml:"HomePhones2"`
	// The ImAddress element contains the primary instant messaging address of a persona.
	ImAddress *ImAddressString `xml:"ImAddress"`
	// The ImAddresses element specifies an array of instant message addresses and the identifiers of their source attributions for the associated persona.
	ImAddresses *ImAddressesArrayOfStringAttributedValuesType `xml:"ImAddresses"`
	// The ImAddresses2 element specifies an array of instant message addresses and the identifiers of their source attributions for the associated persona.
	ImAddresses2 *ImAddresses2 `xml:"ImAddresses2"`
	// The ImAddresses3 element specifies an array of instant message addresses and the identifiers of their source attributions for the associated persona.
	ImAddresses3 *ImAddresses3 `xml:"ImAddresses3"`
	// The Initials element specifies an array of initials values and the identifiers of their source attributions for the associated persona.
	Initials *InitialsArrayOfStringAttributedValuesType `xml:"Initials"`
	// The Location element represents the location of a meeting, appointment, or persona.
	Location *Location `xml:"Location"`
	// The Locations element specifies an array of location values and the identifiers of their source attributions for the associated persona.
	Locations *Locations `xml:"Locations"`
	// The Managers element specifies an array of manager names and the identifiers of their source attributions for a persona.
	Managers *Managers `xml:"Managers"`
	// The MiddleName element represents the middle name of a contact.
	MiddleName *MiddleName `xml:"MiddleName"`
	// The MiddleNames element specifies an array of middle name values and the identifiers of their source attributions for the associated persona.
	MiddleNames *MiddleNames `xml:"MiddleNames"`
	// The MobilePhones element specifies an array of mobile phone numbers and the identifiers of their source attributions for the associated persona.
	MobilePhones *MobilePhones `xml:"MobilePhones"`
	// The MobilePhones2 element specifies an array of MobilePhone values and the identifiers of their source attributions for the associated persona.
	MobilePhones2 *MobilePhones2 `xml:"MobilePhones2"`
	// The Nickname element represents the nickname of a contact.
	Nickname *Nickname `xml:"Nickname"`
	// The Nicknames element specifies an array of nickname values and the identifiers of their source attributions for the associated persona.
	Nicknames *Nicknames `xml:"Nicknames"`
	// The OfficeLocations element specifies an array of office locations and the identifiers of their source attributions for the associated persona.
	OfficeLocations *OfficeLocations `xml:"OfficeLocations"`
	// The OrganizationMainPhones element specifies an array of organizational main phone numbers and the identifiers of their source attributions for the associated persona.
	OrganizationMainPhones *OrganizationMainPhones `xml:"OrganizationMainPhones"`
	// The OtherAddresses element specifies an array of address values and the identifiers of their source attributions for the associated persona.
	OtherAddresses *OtherAddresses `xml:"OtherAddresses"`
	// The OtherFaxes element specifies an array of fax phone number values and the identifiers of their source attributions for the associated persona.
	OtherFaxes *OtherFaxes `xml:"OtherFaxes"`
	// The OtherPhones2 element specifies an array of phone values and the identifiers of their source attributions for the associated persona.
	OtherPhones2 *OtherPhones2 `xml:"OtherPhones2"`
	// The OtherTelephones element specifies an array of telephone values and the identifiers of their source attributions for the associated persona.
	OtherTelephones *OtherTelephones `xml:"OtherTelephones"`
	// The Pagers element specifies an array of pager phone numbers and the identifiers of their source attributions for the associated persona.
	Pagers *Pagers `xml:"Pagers"`
	// The PersonaId element specifies the persona identifier for the associated persona.
	PersonaId *PersonaId `xml:"PersonaId"`
	// The PersonaObjectStatus element specifies whether the information in the associated persona is complete or partial.
	PersonaObjectStatus *PersonaObjectStatus `xml:"PersonaObjectStatus"`
	// The PersonaType element specifies the type of the persona, for example, a person or a distribution list.
	PersonaType *PersonaType `xml:"PersonaType"`
	// The PersonalHomePages element specifies an array of home pages and the identifiers of their source attributions for the associated persona.
	PersonalHomePages *PersonalHomePages `xml:"PersonalHomePages"`
	// The PhoneNumber element specifies the default phone number of the associated persona.
	PhoneNumber *PhoneNumber `xml:"PhoneNumber"`
	// The Professions element specifies an array of Profession values and the identifiers of their source attributions for the associated persona.
	Professions *Professions `xml:"Professions"`
	// The RadioPhones element specifies an array of radio phone numbers and the identifiers of their source attributions for the associated persona.
	RadioPhones *RadioPhones `xml:"RadioPhones"`
	// The RelevanceScore element specifies an integer that represents how relevant the associated persona is to the client.
	RelevanceScore *RelevanceScore `xml:"RelevanceScore"`
	// The Schools element specifies an array of school names and the identifiers of their source attributions for the associated persona.
	Schools *Schools `xml:"Schools"`
	// The SpouseNames element specifies an array of spouse or partner names and the identifiers of their source attributions for the associated persona.
	SpouseNames *SpouseNames `xml:"SpouseNames"`
	// The Surname element represents the surname of a contact.
	Surname *Surname `xml:"Surname"`
	// The Surnames element specifies an array of surname values and the identifiers of their source attributions for the associated persona.
	Surnames *Surnames `xml:"Surnames"`
	// The TTYTDDPhoneNumbers element specifies an array of TTY or TDD text telephone numbers and the identifiers of their source attributions for the associated persona.
	TTYTDDPhoneNumbers *TTYTDDPhoneNumbers `xml:"TTYTDDPhoneNumbers"`
	// The TelexNumbers element specifies an array of Telex numbers and the identifiers of their source attributions for the associated persona.
	TelexNumbers *TelexNumbers `xml:"TelexNumbers"`
	// The Title element represents the title of a contact.
	Title *Title `xml:"Title"`
	// The Titles element specifies an array of job titles and the identifiers of their source attributions for the associated persona.
	Titles *Titles `xml:"Titles"`
	// The WeddingAnniversaries element specifies an array of wedding anniversary dates, stored as strings, and the identifiers of their source attributions for the associated persona.
	WeddingAnniversaries *WeddingAnniversaries `xml:"WeddingAnniversaries"`
	// The WorkCity element specifies the city where the associated persona works.
	WorkCity *WorkCity `xml:"WorkCity"`
	// The WorkCitySortKey element contains the sort key for the work city values associated with a persona.
	WorkCitySortKey *WorkCitySortKey `xml:"WorkCitySortKey"`
	// The WorkFaxes element specifies an array of work fax numbers and the identifiers of their source attributions for the associated persona.
	WorkFaxes *WorkFaxes `xml:"WorkFaxes"`
	// The YomiCompanyName element specifies the phonetic Japanese company name of the associated persona.
	YomiCompanyName *YomiCompanyName `xml:"YomiCompanyName"`
	// The YomiCompanyNames element specifies an array of phonetic Japanese company names and the identifiers of their source attributions for the associated persona.
	YomiCompanyNames *YomiCompanyNames `xml:"YomiCompanyNames"`
	// The YomiFirstName element represents the name that is used in Japan for the searchable or phonetic spelling for a Japanese first name.
	YomiFirstName *YomiFirstName `xml:"YomiFirstName"`
	// The YomiFirstNames element specifies an array of phonetic Japanese first names and the identifiers of their source attributions for the associated persona.
	YomiFirstNames *YomiFirstNames `xml:"YomiFirstNames"`
	// The YomiLastName element represents the name that is used in Japan for the searchable or phonetic spelling for a Japanese last name.
	YomiLastName *YomiLastName `xml:"YomiLastName"`
	// The YomiLastNames element specifies an array of phonetic Japanese last names and the identifiers of their source attributions for the associated persona.
	YomiLastNames *YomiLastNames `xml:"YomiLastNames"`
}

func (P *Persona) SetForMarshal() {
	P.XMLName.Local = "m:Persona"
}

func (P *Persona) GetSchema() *Schema {
	return &SchemaMessages
}
