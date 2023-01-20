package elements

// The CopyItemResponseMessage element contains the status and result of a single CopyItem operation request.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/copyitemresponsemessage
import "encoding/xml"

type CopyItemResponseMessage struct {
	XMLName xml.Name

	// The DescriptiveLinkKey element is currently unused and is reserved for future use. It contains a value of 0.
	DescriptiveLinkKey *DescriptiveLinkKey `xml:"DescriptiveLinkKey"`
	// The Items element contains an array of items.
	Items *Items `xml:"Items"`
	// The MessageText element provides a text description of the status of the response.
	MessageText *MessageText `xml:"MessageText"`
	// The MessageXml element provides additional error response information.
	MessageXml *MessageXml `xml:"MessageXml"`
	// The ResponseCode element provides status information about the request.
	ResponseCode *ResponseCode `xml:"ResponseCode"`
	// Describes the status of a CopyItem operation response.The following values are valid for this attribute:- Success  - Warning  - Error
	ResponseClass *string `xml:"ResponseClass,attr"`
}

const (
	// Describes a request that is fulfilled.
	CopyItemResponseMessageSuccess = `Success`
	// Describes a request that was not processed. A warning may be returned if an error occurred while an item in the request was processing and subsequent items could not be processed.The following are examples of sources of warnings:- The Exchange store goes offline during the batch.  - Active Directory Domain Services (AD DS) goes offline.  - Mailboxes are moved.  - The mailbox database (MDB) goes offline.  - A password is expired.  - A quota is exceeded.
	CopyItemResponseMessageWarning = `Warning`
	// Describes a request that cannot be fulfilled.The following are examples of sources of errors:  - Invalid attributes or elements  - Attributes or elements out of range  - Unknown tag  - Attribute or element not valid in the context  - Unauthorized access attempt by any client  - Server-side failure in response to a valid client-side callInformation about the error can be found in the ResponseCode and MessageText elements.
	CopyItemResponseMessageError = `Error`
)

func (C *CopyItemResponseMessage) SetForMarshal() {
	C.XMLName.Local = "m:CopyItemResponseMessage"
}

func (C *CopyItemResponseMessage) GetSchema() *Schema {
	return &SchemaMessages
}
