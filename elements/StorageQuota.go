package elements

// The StorageQuota element describes the storage quota for the managed folder.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/storagequota
import "encoding/xml"

type StorageQuota struct {
	XMLName xml.Name
	TEXT    int64 `xml:",chardata"`
}

func (S *StorageQuota) SetForMarshal() {
	S.XMLName.Local = "t:StorageQuota"
}

func (S *StorageQuota) GetSchema() *Schema {
	return &SchemaTypes
}
