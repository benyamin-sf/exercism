package speed

// TODO: define the 'Car' type struct
type Car struct {
	battery      int
	batteryDrain int
	speed        int
	distance     int
}

// NewCar creates a new remote controlled car with full battery and given specifications.
func NewCar(speed, batteryDrain int) Car {
	return Car{
		battery:      100,
		speed:        speed,
		batteryDrain: batteryDrain,
		distance:     0,
	}
}

type Track struct {
	distance int
}

// NewTrack creates a new track
func NewTrack(distance int) Track {
	return Track{
		distance: distance,
	}
}

func Drive(car Car) Car {
    if car.battery < car.batteryDrain { return car }
    
	car.battery -= car.batteryDrain
	car.distance += car.speed
	return car
}

// CanFinish checks if a car is able to finish a certain track.
func CanFinish(car Car, track Track) bool {
    if car.battery < car.batteryDrain { return false }

	drivesToFinish := track.distance / car.speed
    totalBatteryNeeded := drivesToFinish * car.batteryDrain

    if car.battery < totalBatteryNeeded { return false }
    return true
}
