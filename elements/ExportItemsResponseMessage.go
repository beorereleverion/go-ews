package elements

// The ExportItemsResponseMessage element contains the status and results of a request to export a single mailbox item.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/exportitemsresponsemessage
import "encoding/xml"

type ExportItemsResponseMessage struct {
	XMLName xml.Name

	// The Data element contains the data of a single exported item or an item to upload into a mailbox.
	Data *Database64Binary `xml:"Data"`
	// The DescriptiveLinkKey element is currently unused and is reserved for future use. It contains a value of 0.
	DescriptiveLinkKey *DescriptiveLinkKey `xml:"DescriptiveLinkKey"`
	// The ItemId element contains the unique identifier and change key of an item in the Exchange store.
	ItemId *ItemId `xml:"ItemId"`
	// The MessageText element provides a text description of the status of the response.
	MessageText *MessageText `xml:"MessageText"`
	// The MessageXml element provides additional error response information.
	MessageXml *MessageXml `xml:"MessageXml"`
	// The ResponseCode element provides status information about the request.
	ResponseCode *ResponseCode `xml:"ResponseCode"`
	// Describes the status of the response. The following values are valid for this attribute:  - Success  - Warning  - Error
	ResponseClass *string `xml:"ResponseClass,attr"`
}

const (
	// Describes a request that is fulfilled.
	ExportItemsResponseMessageSuccess = `Success`
	// Describes a request that was not processed. A warning may be returned if an error occurred while an item in the request was processing and subsequent items could not be processed. The following are examples of sources of warnings:  - The Exchange store is offline during the batch.  - Active Directory Domain Services (AD DS) is offline.  - Mailboxes were moved.  - The message database (MDB) is offline.  - A password is expired.  - A quota has been exceeded.
	ExportItemsResponseMessageWarning = `Warning`
	// Describes a request that cannot be fulfilled. The following are examples of sources of errors:  - Invalid attributes or elements  - Attributes or elements that are out of range  - An unknown tag  - An attribute or element that is not valid in the context  - An unauthorized access attempt by any client  - A server-side failure in response to a valid client-side call    Information about the error can be found in the ResponseCode and MessageText elements.
	ExportItemsResponseMessageError = `Error`
)

func (E *ExportItemsResponseMessage) SetForMarshal() {
	E.XMLName.Local = "m:ExportItemsResponseMessage"
}

func (E *ExportItemsResponseMessage) GetSchema() *Schema {
	return &SchemaMessages
}
