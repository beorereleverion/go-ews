package elements

// The SendNotificationResult element contains the response of a client application to a push notification.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/sendnotificationresult
import "encoding/xml"

type SendNotificationResult struct {
	XMLName xml.Name

	// The SubscriptionStatus element describes the status of a push subscription.
	SubscriptionStatus *SubscriptionStatus `xml:"SubscriptionStatus"`
}

func (S *SendNotificationResult) SetForMarshal() {
	S.XMLName.Local = "m:SendNotificationResult"
}

func (S *SendNotificationResult) GetSchema() *Schema {
	return &SchemaMessages
}
