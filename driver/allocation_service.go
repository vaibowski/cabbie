package driver

import (
	"cabbie/models"
	"errors"
	"github.com/emirpasic/gods/v2/maps/treemap"
	"log"
	"slices"
)

type AllocationService struct {
	// we will have 4 maps, each corresponding to a unique driverService type
	ActiveDriverPool []*treemap.Map[float64, []string]
}

func (svc *AllocationService) AllocateDriver(pickup models.Location, serviceType models.ServiceTypeEnum) (string, error) {
	driverPoolForServiceType := svc.ActiveDriverPool[serviceType]

	if driverPoolForServiceType.Size() == 0 {
		return "", errors.New("no available locations")
	}

	floor, floorDriverList, foundFloor := driverPoolForServiceType.Floor(pickup.XCoordinate)
	ceiling, ceilingDriverList, foundCeiling := driverPoolForServiceType.Ceiling(pickup.XCoordinate)

	var driverIDList []string
	var targetLocation float64
	if !foundFloor {
		driverIDList = ceilingDriverList
		targetLocation = ceiling
	} else if !foundCeiling {
		driverIDList = floorDriverList
		targetLocation = floor
	} else {
		if pickup.XCoordinate-floor <= ceiling-pickup.XCoordinate {
			driverIDList = floorDriverList
			targetLocation = floor
		} else {
			driverIDList = ceilingDriverList
			targetLocation = ceiling
		}
	}

	// target location is found, now we will just assign the first driver we find, and remove them from the treemap
	driverID := driverIDList[0]
	driverIDList = slices.Delete(driverIDList, 0, 1)
	if len(driverIDList) == 0 {
		svc.ActiveDriverPool[serviceType].Remove(targetLocation)
	} else {
		svc.ActiveDriverPool[serviceType].Put(targetLocation, driverIDList)
	}

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
