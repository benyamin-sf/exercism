package purchase

import "fmt"

// NeedsLicense determines whether a license is needed to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {
	return kind == "car" || kind == "truck"
}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in lexicographical order.
func ChooseVehicle(option1, option2 string) string {
	preferredCar := option1
    if preferredCar > option2 {
        preferredCar = option2
    }
    return fmt.Sprintf("%s is clearly the better choice.", preferredCar)
}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {
	var finalPrice float64
    
	if age < 3 {
        finalPrice = originalPrice * 0.8
    } else if age >= 3 && age < 10 {
        finalPrice = originalPrice * 0.7
    } else {
        finalPrice = originalPrice * 0.5
    }

    return finalPrice
}
