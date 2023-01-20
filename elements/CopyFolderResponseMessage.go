package elements

// The CopyFolderResponseMessage element contains the status and result of a single CopyFolder operation request.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/copyfolderresponsemessage
import "encoding/xml"

type CopyFolderResponseMessage struct {
	XMLName xml.Name

	// The DescriptiveLinkKey element is currently unused and is reserved for future use. It contains a value of 0.
	DescriptiveLinkKey *DescriptiveLinkKey `xml:"DescriptiveLinkKey"`
	// The Folders element contains an array of folders that are used in folder operations.
	Folders *Folders `xml:"Folders"`
	// The MessageText element provides a text description of the status of the response.
	MessageText *MessageText `xml:"MessageText"`
	// The MessageXml element provides additional error response information.
	MessageXml *MessageXml `xml:"MessageXml"`
	// The ResponseCode element provides status information about the request.
	ResponseCode *ResponseCode `xml:"ResponseCode"`
	// Describes the status of a CopyFolder operation response.The following values are valid for this attribute:- Success  - Warning  - Error
	ResponseClass *string `xml:"ResponseClass,attr"`
}

const (
	// Describes a request that is fulfilled.
	CopyFolderResponseMessageSuccess = `Success`
	// Describes a request that was not processed. A warning may be returned if an error occurred while an item in the request was processing and subsequent items could not be processed.The following are examples of sources of warnings:- The Exchange store goes offline during the batch.  - Active Directory Domain Services (AD DS) goes offline.  - Mailboxes are moved.  - The message database (MDB) goes offline.  - A password is expired.  - A quota is exceeded.
	CopyFolderResponseMessageWarning = `Warning`
	// Describes a request that cannot be fulfilled.The following are examples of sources of errors:  - Invalid attributes or elements  - Attributes or elements out of range  - Unknown tag  - Attribute or element not valid in the context  - Unauthorized access attempted by any client  - Server-side failure in response to a valid client-side callInformation about the error can be found in the ResponseCode and MessageText elements.
	CopyFolderResponseMessageError = `Error`
)

func (C *CopyFolderResponseMessage) SetForMarshal() {
	C.XMLName.Local = "m:CopyFolderResponseMessage"
}

func (C *CopyFolderResponseMessage) GetSchema() *Schema {
	return &SchemaMessages
}
