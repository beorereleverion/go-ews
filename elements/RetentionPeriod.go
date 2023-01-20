package elements

// The RetentionPeriod element specifies the number of days that the retention policy is in effect.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/retentionperiod
import "encoding/xml"

type RetentionPeriod struct {
	XMLName xml.Name
	TEXT    int64 `xml:",chardata"`
}

func (R *RetentionPeriod) SetForMarshal() {
	R.XMLName.Local = "t:RetentionPeriod"
}

func (R *RetentionPeriod) GetSchema() *Schema {
	return &SchemaTypes
}
