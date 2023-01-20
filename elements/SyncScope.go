package elements

// The SyncScope element specifies whether just items or items and folder associated information are returned in a synchronization response.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/syncscope
import "encoding/xml"

type SyncScope struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

const (
	// NormalAndAssociatedItems
	SyncScopeNormalAndAssociatedItems string = `NormalAndAssociatedItems`
	// NormalItems
	SyncScopeNormalItems string = `NormalItems`
)

func (S *SyncScope) SetForMarshal() {
	S.XMLName.Local = "m:SyncScope"
}

func (S *SyncScope) GetSchema() *Schema {
	return &SchemaMessages
}
