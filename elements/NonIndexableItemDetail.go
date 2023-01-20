package elements

// The NonIndexableItemDetail element specifies detail information about an item that cannot be indexed.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/nonindexableitemdetail
import "encoding/xml"

type NonIndexableItemDetail struct {
	XMLName xml.Name

	// The AdditionalInfo element specifies additional information about the hold status of a mailbox.
	AdditionalInfo *AdditionalInfo `xml:"AdditionalInfo"`
	// The AttemptCount element represents the number of attempts that have been made to index the item.
	AttemptCount *AttemptCount `xml:"AttemptCount"`
	// The ErrorCode element is intended for internal use only.
	ErrorCode *ErrorCodeItemIndexErrorType `xml:"ErrorCode"`
	// The ErrorDescription element describes the error that is returned in information about an item that cannot be indexed.
	ErrorDescription *ErrorDescription `xml:"ErrorDescription"`
	// The IsPartiallyIndexed element indicates whether the item is partially indexed.
	IsPartiallyIndexed *IsPartiallyIndexed `xml:"IsPartiallyIndexed"`
	// The IsPermanentFailure element indicates whether a previous attempt to index the item was unsuccessful.
	IsPermanentFailure *IsPermanentFailure `xml:"IsPermanentFailure"`
	// The ItemId element contains the unique identifier and change key of an item in the Exchange store.
	ItemId *ItemId `xml:"ItemId"`
	// The LastAttemptTime element contains the time and date at which the last attempt to index the item was made.
	LastAttemptTime *LastAttemptTime `xml:"LastAttemptTime"`
	// The SortValue element specifies a value used for sorting.
	SortValue *SortValue `xml:"SortValue"`
}

func (N *NonIndexableItemDetail) SetForMarshal() {
	N.XMLName.Local = "t:NonIndexableItemDetail"
}

func (N *NonIndexableItemDetail) GetSchema() *Schema {
	return &SchemaTypes
}
