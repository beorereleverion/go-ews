package elements

// The GetSharingFolder element defines a request to get the local folder identifier of a specified shared folder. It is the base element for the GetSharingFolder operation.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/getsharingfolder
type GetSharingFolder struct {
	// The DataType element describes the type of data that is shared by a shared folder.
	DataType *DataType `xml:"m:DataType"`
	// The SharedFolderId element represents the identifier of the shared folder the local folder identifier for which should be returned by a GetSharingFolder operation request.
	SharedFolderId *SharedFolderId `xml:"m:SharedFolderId"`
	// The SmtpAddress element represents the Simple Mail Transfer Protocol (SMTP) address of an account to be used for impersonation or a Simple Mail Transfer Protocol (SMTP) recipient address of a calendar or contact sharing request.
	SmtpAddress *SmtpAddress `xml:"t:SmtpAddress"`
}
