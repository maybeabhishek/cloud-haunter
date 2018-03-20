package operation

import (
	log "github.com/Sirupsen/logrus"
	"github.com/hortonworks/cloud-cost-reducer/context"
	"github.com/hortonworks/cloud-cost-reducer/types"
)

func init() {
	context.Operations[types.OWNERLESS] = Ownerless{}
}

type Ownerless struct {
}

func (o Ownerless) Execute(clouds []types.CloudType) []*types.Instance {
	allInstances := []*types.Instance{}
	for _, cloud := range clouds {
		provider, ok := context.CloudProviders[cloud]
		if !ok {
			panic("Cloud provider not supported")
		}
		instances, err := provider.GetOwnerLessInstances()
		if err != nil {
			log.Errorf("[OWNERLESS] Failed to collect owner less instances on %s, err: %s", cloud.String(), err.Error())
			continue
		}
		allInstances = append(allInstances, instances...)
	}
	for _, instance := range allInstances {
		log.Infof("[%s] Instance %s is running without owner tag", instance.CloudType, instance.Name)
	}
	return allInstances
}
