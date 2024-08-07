package driver

import (
	"cabbie/models"
	"errors"
	"github.com/emirpasic/gods/v2/maps/treemap"
	"log"
	"math"
	"slices"
)

type AllocationService struct {
	// we will have 4 maps, each corresponding to a unique driverService type
	ActiveDriverPool []*treemap.Map[float64, []string]
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

func (svc *AllocationService) SetLocation(driverID string, serviceType models.ServiceTypeEnum, location models.Location) {
	driversAtLocation, found := svc.ActiveDriverPool[serviceType].Get(location.XCoordinate)
	if found != true || driversAtLocation == nil {
		var driverList []string
		driverList = append(driverList, driverID)
		svc.ActiveDriverPool[serviceType].Put(location.XCoordinate, driverList)
	} else {
		driversAtLocation = append(driversAtLocation, driverID)
		svc.ActiveDriverPool[serviceType].Put(location.XCoordinate, driversAtLocation)
	}
	return
}

func (svc *AllocationService) UnsetLocation(driverID string, serviceType models.ServiceTypeEnum, lastLocation models.Location) error {
	driverListAtLocation, found := svc.ActiveDriverPool[serviceType].Get(lastLocation.XCoordinate)
	if found != true || driverListAtLocation == nil {
		log.Printf("no drivers listed at this location: %f", lastLocation)
		return errors.New("no drivers listed at this location")
	} else {
		driverIndex := slices.Index(driverListAtLocation, driverID)
		if driverIndex == -1 {
			log.Printf("driverID %s not found at previous location %f", driverID, lastLocation)
			return errors.New("driver not found at previous location")
		} else {
			driverListAtLocation = slices.Delete(driverListAtLocation, driverIndex, driverIndex+1)
			svc.ActiveDriverPool[serviceType].Put(lastLocation.XCoordinate, driverListAtLocation)
		}
	}
	return nil
}

func (svc *AllocationService) GetActiveDriverPool() map[int64]map[float64][]string {
	result := make(map[int64]map[float64][]string)
	for i := 1; i <= 4; i++ {
		driverPool := svc.ActiveDriverPool[i]
		locationDriverMap := make(map[float64][]string)
		for _, location := range driverPool.Keys() {
			driverList, _ := driverPool.Get(location)
			locationDriverMap[location] = driverList
		}
		result[int64(i)] = locationDriverMap
	}
	return result
}

func NewAllocationService(activeDriverPool []*treemap.Map[float64, []string]) AllocationService {
	return AllocationService{
		ActiveDriverPool: activeDriverPool,
	}
}
