package operation

import (
	log "github.com/Sirupsen/logrus"
	ctx "github.com/hortonworks/cloud-haunter/context"
	"github.com/hortonworks/cloud-haunter/types"
)

func init() {
	ctx.Operations[types.Databases] = databases{}
}

type databases struct {
}

func (o databases) Execute(clouds []types.CloudType) []types.CloudItem {
	log.Debugf("[GETDATABASES] Collecting databases on: [%s]", clouds)
	itemsChan, errChan := o.collect(clouds)
	return wait(itemsChan, errChan, "[GETDATABASES] Failed to collect databases")
}

func (o databases) collect(clouds []types.CloudType) (chan []types.CloudItem, chan error) {
	return collect(clouds, func(provider types.CloudProvider) ([]types.CloudItem, error) {
		databases, err := provider.GetDatabases()
		if err != nil {
			return nil, err
		}
		return o.convertToCloudItems(databases), nil
	})
}

func (o databases) convertToCloudItems(databases []*types.Database) []types.CloudItem {
	var items []types.CloudItem
	for _, access := range databases {
		items = append(items, access)
	}
	return items
}
