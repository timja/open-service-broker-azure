package cosmosdb

type commonRegisteredManager struct {
	cosmosAccountManager
}

type sqlAllInOneRegisteredManager struct {
	sqlAllInOneManager
}

type mongoAccountRegisteredManager struct {
	commonRegisteredManager
}
