package elements

// The RetentionAction element specifies the action performed on items with the retention tag.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/retentionaction
import "encoding/xml"

type RetentionAction struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

const (
	// DeleteAndAllowRecovery
	RetentionActionDeleteAndAllowRecovery string = `DeleteAndAllowRecovery`
	// MarkAsPastRetentionLimit
	RetentionActionMarkAsPastRetentionLimit string = `MarkAsPastRetentionLimit`
	// MoveToArchive
	RetentionActionMoveToArchive string = `MoveToArchive`
	// MoveToDeletedItems
	RetentionActionMoveToDeletedItems string = `MoveToDeletedItems`
	// MoveToFolder
	RetentionActionMoveToFolder string = `MoveToFolder`
	// None
	RetentionActionNone string = `None`
	// PermanentlyDelete
	RetentionActionPermanentlyDelete string = `PermanentlyDelete`
)

func (R *RetentionAction) SetForMarshal() {
	R.XMLName.Local = "t:RetentionAction"
}

func (R *RetentionAction) GetSchema() *Schema {
	return &SchemaTypes
}
