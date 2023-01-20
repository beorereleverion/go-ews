package elements

// The SyncState element contains a base64-encoded form of the synchronization data that is updated after each successful request. This is used to identify the synchronization state.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/syncstate-ex15websvcsotherref
import "encoding/xml"

type SyncState struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (S *SyncState) SetForMarshal() {
	S.XMLName.Local = "m:SyncState"
}

func (S *SyncState) GetSchema() *Schema {
	return &SchemaMessages
}
