package driver

import (
	"cabbie/models"
	"errors"
	"github.com/emirpasic/gods/v2/maps/treemap"
	"math"
	"slices"
)

type AllocationService struct {
	// we will have 4 maps, each corresponding to a unique service type
	ActiveDriverPool []*treemap.Map[float64, []string]
}

func NewAllocationService(activeDriverPool []*treemap.Map[float64, []string]) AllocationService {
	return AllocationService{
		ActiveDriverPool: activeDriverPool,
	}
}

func (svc *AllocationService) AllocateDriver(pickup models.Location, serviceType models.ServiceTypeEnum) (string, error) {
	keys := svc.ActiveDriverPool[serviceType].Keys()
	size := len(keys)
	if size == 0 {
		return "", errors.New("no available locations")
	}
	var target float64
	var minDistance float64
	minDistance = 1000

	// iterate
	for i := range keys {
		if math.Abs(pickup.XCoordinate-keys[i]) < minDistance && svc.ActiveDriverPool[serviceType].Size() != 0 {
			target = keys[i]
		}
	}

	// target location is found, now we will just assign the first driver we find, and remove them from the treemap
	driverIDList, _ := svc.ActiveDriverPool[serviceType].Get(target)
	driverID := driverIDList[0]
	driverIDList = slices.Delete(driverIDList, 0, 1)
	svc.ActiveDriverPool[serviceType].Put(target, driverIDList)
	return driverID, nil
}
