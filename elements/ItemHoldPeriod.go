package elements

// The ItemHoldPeriod element specifies the amount of time to hold content that matches the mailbox query.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/itemholdperiod
import "encoding/xml"

type ItemHoldPeriod struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (I *ItemHoldPeriod) SetForMarshal() {
	I.XMLName.Local = "m:ItemHoldPeriod"
}

func (I *ItemHoldPeriod) GetSchema() *Schema {
	return &SchemaMessages
}
