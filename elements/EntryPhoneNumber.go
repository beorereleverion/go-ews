package elements

// The Entry element represents a telephone number for a contact.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/entry-phonenumber
type EntryPhoneNumber struct {
	// Identifies the telephone number. The Key attribute is of type PhoneNumberKeyType. The following are the possible values for this attribute:- AssistantPhone  - BusinessFax  - BusinessPhone  - BusinessPhone2  - Callback  - CarPhone  - CompanyMainPhone  - HomeFax  - HomePhone  - HomePhone2  - Isdn  - MobilePhone  - OtherFax  - OtherTelephone  - Pager  - PrimaryPhone  - RadioPhone  - Telex  - TtyTddPhone
	Key *string `xml:"Key,attr"`
}
