package elements

// The SyncFolderHierarchyResponseMessage element contains the status and result of a single SyncFolderHierarchy operation request.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/syncfolderhierarchyresponsemessage
import "encoding/xml"

type SyncFolderHierarchyResponseMessage struct {
	XMLName xml.Name

	// The Changes element contains a sequenced array of change types that represent the type of differences between the folders on the client and the folders on the computer that is running Microsoft Exchange Server 2007.
	Changes *ChangesHierarchy `xml:"Changes"`
	// The DescriptiveLinkKey element is currently unused and is reserved for future use. It contains a value of 0.
	DescriptiveLinkKey *DescriptiveLinkKey `xml:"DescriptiveLinkKey"`
	// The IncludesLastFolderInRange element indicates whether the last item to synchronize has been included in the response.
	IncludesLastFolderInRange *IncludesLastFolderInRange `xml:"IncludesLastFolderInRange"`
	// The MessageText element provides a text description of the status of the response.
	MessageText *MessageText `xml:"MessageText"`
	// The MessageXml element provides additional error response information.
	MessageXml *MessageXml `xml:"MessageXml"`
	// The ResponseCode element provides status information about the request.
	ResponseCode *ResponseCode `xml:"ResponseCode"`
	// The SyncState element contains a base64-encoded form of the synchronization data that is updated after each successful request. This is used to identify the synchronization state.
	SyncState *SyncState `xml:"SyncState"`
	// Describes the status of a SyncFolderHierarchy operation response. The following values are valid for this attribute:  - Success  - Warning  - Error
	ResponseClass *string `xml:"ResponseClass,attr"`
}

const (
	// Describes a request that is fulfilled.
	SyncFolderHierarchyResponseMessageSuccess = `Success`
	// Describes a request that was not processed. A warning may be returned if an error occurred while an item in the request was processing and subsequent items could not be processed. The following are examples of sources of warnings:  - The Exchange store is offline during the batch.  - Active Directory Domain Services (AD DS) is offline.  - Mailboxes were moved.  - The message database (MDB) is offline.  - A password has expired.  - A quota has been exceeded.
	SyncFolderHierarchyResponseMessageWarning = `Warning`
	// Describes a request that cannot be fulfilled. The following are examples of sources of errors:  - Invalid attributes or elements  - Attributes or elements that are out of range  - An unknown tag  - An attribute or element that is not valid in the context  - An unauthorized access attempt by any client  - A server-side failure in response to a valid client-side call    Information about the error can be found in the ResponseCode and MessageText elements.
	SyncFolderHierarchyResponseMessageError = `Error`
)

func (S *SyncFolderHierarchyResponseMessage) SetForMarshal() {
	S.XMLName.Local = "m:SyncFolderHierarchyResponseMessage"
}

func (S *SyncFolderHierarchyResponseMessage) GetSchema() *Schema {
	return &SchemaMessages
}
