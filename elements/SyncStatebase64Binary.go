package elements

// The SyncState element specifies the synchronization state of a conversation.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/syncstate-base64binary
import "encoding/xml"

type SyncStatebase64Binary struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (S *SyncStatebase64Binary) SetForMarshal() {
	S.XMLName.Local = "t:SyncState"
}

func (S *SyncStatebase64Binary) GetSchema() *Schema {
	return &SchemaTypes
}
