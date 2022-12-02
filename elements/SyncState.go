package elements

// The SyncState element contains a base64-encoded form of the synchronization data that is updated after each successful request. This is used to identify the synchronization state.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/syncstate-ex15websvcsotherref
type SyncState string
