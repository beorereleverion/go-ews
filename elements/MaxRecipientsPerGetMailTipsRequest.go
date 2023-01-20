package elements

// The MaxRecipientsPerGetMailTipsRequest element indicates the maximum number of recipients that can be passed to the GetMailTips operation.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/maxrecipientspergetmailtipsrequest
import "encoding/xml"

type MaxRecipientsPerGetMailTipsRequest struct {
	XMLName xml.Name
	TEXT    int64 `xml:",chardata"`
}

func (M *MaxRecipientsPerGetMailTipsRequest) SetForMarshal() {
	M.XMLName.Local = "t:MaxRecipientsPerGetMailTipsRequest"
}

func (M *MaxRecipientsPerGetMailTipsRequest) GetSchema() *Schema {
	return &SchemaTypes
}
