package elements

// The ImListMigrationCompleted element indicates whether the Exchange store contains the instant messaging items used by instant messaging clients.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/imlistmigrationcompleted
type ImListMigrationCompleted bool

const (
	// false
	ImListMigrationCompletedfalse ImListMigrationCompleted = false
	// true
	ImListMigrationCompletedtrue ImListMigrationCompleted = true
)
