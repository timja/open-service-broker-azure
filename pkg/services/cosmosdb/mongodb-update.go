package cosmosdb

import (
	"github.com/Azure/open-service-broker-azure/pkg/service"
)

func (
	m *mongoAccountManager,
) ValidateUpdatingParameters(instance service.Instance) error {
	return readRegionsValidator(
		"mongo account update",
		[]interface{}{instance.UpdatingParameters.GetStringArray("readLocations")},
	)
}

func (
	m *mongoAccountManager,
) GetUpdater(service.Plan) (service.Updater, error) {
	return service.NewUpdater(
		service.NewUpdatingStep("updateReadRegions", m.updateReadRegions),
	)
}
