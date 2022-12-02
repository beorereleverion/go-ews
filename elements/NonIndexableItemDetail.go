package elements

// The NonIndexableItemDetail element specifies detail information about an item that cannot be indexed.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/nonindexableitemdetail
type NonIndexableItemDetail struct {
	// The AdditionalInfo element specifies additional information about the hold status of a mailbox.
	AdditionalInfo *AdditionalInfo `xml:"t:AdditionalInfo"`
	// The AttemptCount element represents the number of attempts that have been made to index the item.
	AttemptCount *AttemptCount `xml:"t:AttemptCount"`
	// The ErrorCode element is intended for internal use only.
	ErrorCode *ErrorCodeItemIndexErrorType `xml:"t:ErrorCode"`
	// The ErrorDescription element describes the error that is returned in information about an item that cannot be indexed.
	ErrorDescription *ErrorDescription `xml:"t:ErrorDescription"`
	// The IsPartiallyIndexed element indicates whether the item is partially indexed.
	IsPartiallyIndexed *IsPartiallyIndexed `xml:"t:IsPartiallyIndexed"`
	// The IsPermanentFailure element indicates whether a previous attempt to index the item was unsuccessful.
	IsPermanentFailure *IsPermanentFailure `xml:"t:IsPermanentFailure"`
	// The ItemId element contains the unique identifier and change key of an item in the Exchange store.
	ItemId *ItemId `xml:"t:ItemId"`
	// The LastAttemptTime element contains the time and date at which the last attempt to index the item was made.
	LastAttemptTime *LastAttemptTime `xml:"t:LastAttemptTime"`
	// The SortValue element specifies a value used for sorting.
	SortValue *SortValue `xml:"t:SortValue"`
}
