package elements

// The SyncFolderItemsResponseMessage element contains the status and result of a single SyncFolderItems operation request.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/syncfolderitemsresponsemessage
import "encoding/xml"

type SyncFolderItemsResponseMessage struct {
	XMLName xml.Name

	// The Changes element contains a sequence array of change types that represent the types of differences between the items on the client and the items on the Exchange server.
	Changes *ChangesItems `xml:"Changes"`
	// The DescriptiveLinkKey element is currently unused and is reserved for future use. It contains a value of 0.
	DescriptiveLinkKey *DescriptiveLinkKey `xml:"DescriptiveLinkKey"`
	// The IncludesLastItemInRange element indicates whether the last item to synchronize has been included in the response.
	IncludesLastItemInRange *IncludesLastItemInRange `xml:"IncludesLastItemInRange"`
	// The MessageText element provides a text description of the status of the response.
	MessageText *MessageText `xml:"MessageText"`
	// The MessageXml element provides additional error response information.
	MessageXml *MessageXml `xml:"MessageXml"`
	// The ResponseCode element provides status information about the request.
	ResponseCode *ResponseCode `xml:"ResponseCode"`
	// The SyncState element contains a base64-encoded form of the synchronization data that is updated after each successful request. This is used to identify the synchronization state.
	SyncState *SyncState `xml:"SyncState"`
	// Describes the status of a SyncFolderItems operation response. The following values are valid for this attribute:  - Success  - Warning  - Error
	ResponseClass *string `xml:"ResponseClass,attr"`
}

func (S *SyncFolderItemsResponseMessage) SetForMarshal() {
	S.XMLName.Local = "m:SyncFolderItemsResponseMessage"
}

func (S *SyncFolderItemsResponseMessage) GetSchema() *Schema {
	return &SchemaMessages
}
