package elements

// The HasQuota element indicates whether the managed folder has a quota.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/hasquota
import "encoding/xml"

type HasQuota struct {
	XMLName xml.Name
	TEXT    bool `xml:",chardata"`
}

func (H *HasQuota) SetForMarshal() {
	H.XMLName.Local = "t:HasQuota"
}

func (H *HasQuota) GetSchema() *Schema {
	return &SchemaTypes
}
