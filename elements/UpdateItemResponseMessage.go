package elements

// The UpdateItemResponseMessage element contains the status and result of a single UpdateItem operation request.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/updateitemresponsemessage
import "encoding/xml"

type UpdateItemResponseMessage struct {
	XMLName xml.Name

	// The ConflictResults element contains the number of conflicts in an UpdateItem operation response.
	ConflictResults *ConflictResults `xml:"ConflictResults"`
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
	// Describes the status of an UpdateItem operation response. The following values are valid for this attribute:  - Success  - Warning  - Error
	ResponseClass *string `xml:"ResponseClass,attr"`
}

const (
	// Describes a request that is fulfilled.
	UpdateItemResponseMessageSuccess = `Success`
	// Describes a request that was not processed. A warning may be returned if an error occurred while an item in the request was processed and subsequent items could not be processed. The following are examples of sources of warnings:  - The Exchange store is offline during the batch.  - Active Directory Domain Services (AD DS) is offline.  - Mailboxes are moved.  - The message database (MDB) is offline.  - A password is expired.  - A quota is exceeded.
	UpdateItemResponseMessageWarning = `Warning`
	// Describes a request that cannot be fulfilled. The following are examples of sources of errors:  - Invalid attributes or elements  - Attributes or elements out of range  - Unknown tag  - Attribute or element not valid in the context  - Unauthorized access attempt by any client  - Server-side failure in response to a valid client-side call    Information about the error can be found in the ResponseCode and MessageText elements.
	UpdateItemResponseMessageError = `Error`
)

func (U *UpdateItemResponseMessage) SetForMarshal() {
	U.XMLName.Local = "m:UpdateItemResponseMessage"
}

func (U *UpdateItemResponseMessage) GetSchema() *Schema {
	return &SchemaMessages
}
